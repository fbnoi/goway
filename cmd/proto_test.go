package main

import (
	pb "flynoob/goway/protobuf"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
)

func BenchmarkMarshal(b *testing.B) {
	heartBeat := &pb.HeartBeat{UpTimestamp: time.Now().UnixMilli()}
	heartBeat.DownTimestamp = time.Now().UnixMilli()
	bs, _ := proto.Marshal(heartBeat)
	newHeartBeat := &pb.HeartBeat{}
	proto.Unmarshal(bs, newHeartBeat)
}
