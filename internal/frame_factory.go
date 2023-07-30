package internal

import (
	pb "flynoob/goway/protobuf"
	"sync"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

var (
	framePool     = sync.Pool{New: func() any { return &pb.Frame{} }}
	heartbeatPool = sync.Pool{New: func() any { return &pb.Heartbeat{} }}
)

func GetFrame(bs []byte) (*pb.Frame, error) {
	frame := framePool.Get().(*pb.Frame)
	if err := proto.Unmarshal(bs, frame); err != nil {
		PutFrame(frame)
		return nil, err
	}

	return frame, nil
}

func PutFrame(frame *pb.Frame) {
	framePool.Put(frame)
}

func GetHearBeatFrame(frame *pb.Frame) (*pb.Heartbeat, error) {
	if frame.Type != pb.FrameType_HEARTBEAT {
		return nil, errors.New("Frame is not a Heartbeat frame")
	}

	heartbeat := heartbeatPool.Get().(*pb.Heartbeat)
	if err := proto.Unmarshal(frame.Body, heartbeat); err != nil {
		return nil, err
	}

	return heartbeat, nil
}

func PutHeartBeat(heartbeat *pb.Heartbeat) {
	heartbeatPool.Put(heartbeat)
}
