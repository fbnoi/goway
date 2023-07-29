package main

import (
	pb "flynoob/goway/protobuf"
	"sync"
	"testing"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func BenchmarkMarshal(b *testing.B) {
	heartBeat := &pb.Heartbeat{UpTimestamp: time.Now().UnixMilli()}
	heartBeat.DownTimestamp = time.Now().UnixMilli()
	bs, _ := proto.Marshal(heartBeat)
	newHeartBeat := &pb.Heartbeat{}
	proto.Unmarshal(bs, newHeartBeat)
}

func TestJob(t *testing.T) {
	s := gocron.NewScheduler(time.UTC)
	wg := sync.WaitGroup{}
	i := 1
	wg.Add(3)
	_, err := s.Every(1).Second().Do(func() {
		i += 1
		t.Log("do job ......")
		wg.Done()
	})
	assert.Equal(t, err, nil)
	s.StartAsync()
	wg.Wait()
	assert.Equal(t, 4, i)
}
