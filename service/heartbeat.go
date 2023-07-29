package service

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

func doOnce() {
	schedulerOnce.Do(func() {
		healthyScanScheduler.StartAsync()
	})
}

func HealthyScan(client *gw.Client) {
	doOnce()
	healthyScanScheduler.Every(1).Second().Do(func() {
		checkHealth(client)
	})
}

func OnPing(client *gw.Client, frame *pb.Frame) error {
	heartbeat, err := GetHearBeatFrame(frame)
	if err != nil {
		return err
	}
	defer PutHeartBeat(heartbeat)
	client.LastPingAt = time.Now()
	heartbeat.DownTimestamp = client.LastPingAt.UnixMilli()
	if frame.Body, err = proto.Marshal(heartbeat); err != nil {
		return err
	}
	client.SendFrame(frame)
	return nil
}

func checkHealth(client *gw.Client) {
	if client.LastPingAt.Add(time.Second).Before(time.Now()) {
		client.Color += 1
	}

	if client.Color == gw.Red {
		client.Close()
	}
}
