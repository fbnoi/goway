package internal

import (
	gw "flynoob/goway"
	pb "flynoob/goway/protobuf"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"google.golang.org/protobuf/proto"
)

var (
	healthyScanScheduler = gocron.NewScheduler(time.UTC)
	schedulerOnce        = sync.Once{}
)

func startSchedule() {
	schedulerOnce.Do(func() {
		healthyScanScheduler.StartAsync()
	})
}

func MonitorHealth(client *gw.Client) {
	startSchedule()
	client.Subscribe(&pb.Heartbeat{}, WrapHandler(client, func(c *gw.Client, m proto.Message) {
		heartbeat, ok := m.(*pb.Heartbeat)
		if !ok {
			return
		}
		client.LastPingAt = time.Now()
		heartbeat.DownTimestamp = client.LastPingAt.UnixMilli()
		client.Send(heartbeat)
	}))
	healthyScanScheduler.Every(1).Second().Do(func() {
		checkHealth(client)
	})
}

func HandleBytesMessage(c *gw.Client, bs []byte) {

}

func WrapHandler(c *gw.Client, fn func(*gw.Client, proto.Message)) gw.HandleFunc {
	return func(m proto.Message) { fn(c, m) }
}

func checkHealth(client *gw.Client) {
	if client.LastPingAt.Add(2 * time.Second).Before(time.Now()) {
		client.Color += 1
	}

	if client.Color == gw.Red {
		client.Close()
	}
}
