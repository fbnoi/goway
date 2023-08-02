package internal

// import (
// 	pb "flynoob/goway/protobuf"
// 	"sync"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestNew(t *testing.T) {
// 	bus := NewBus()
// 	if bus == nil {
// 		t.Log("New EventBus not created!")
// 		t.Fail()
// 	}
// }

// func TestHasCallback(t *testing.T) {
// 	bus := NewBus()
// 	bus.Subscribe(pb.FrameType_CLOSE, func(*pb.Frame) {})
// 	assert.Equal(t, true, bus.HasCallback(pb.FrameType_CLOSE))
// 	assert.Equal(t, false, bus.HasCallback(pb.FrameType_HEARTBEAT))
// }

// func TestSubscribe(t *testing.T) {
// 	bus := NewBus()
// 	bus.Subscribe(pb.FrameType_CLOSE, func(*pb.Frame) {})
// 	assert.Equal(t, true, bus.HasCallback(pb.FrameType_CLOSE))
// }

// func TestSubscribeOnce(t *testing.T) {
// 	bus := NewBus()
// 	bus.SubscribeOnce(pb.FrameType_CLOSE, func(*pb.Frame) {})
// 	assert.Equal(t, true, bus.HasCallback(pb.FrameType_CLOSE))
// }

// func TestSubscribeOnceAndManySubscribe(t *testing.T) {
// 	bus := NewBus()
// 	event := pb.FrameType_CLOSE
// 	flag := 0
// 	fn := func(*pb.Frame) {
// 		flag += 1
// 	}
// 	bus.SubscribeOnce(event, fn)
// 	bus.Subscribe(event, fn)
// 	bus.Subscribe(event, fn)
// 	bus.Publish(&pb.Frame{Type: event})
// 	assert.Equal(t, 3, flag)
// }

// func TestUnsubscribe(t *testing.T) {
// 	bus := NewBus()
// 	handler := func(*pb.Frame) {}
// 	bus.Subscribe(pb.FrameType_CLOSE, handler)
// 	assert.Nil(t, bus.Unsubscribe(pb.FrameType_CLOSE, handler))
// 	assert.NotNil(t, bus.Unsubscribe(pb.FrameType_CLOSE, handler))
// }

// type handler struct {
// 	val int
// }

// func (h *handler) Handle(*pb.Frame) {
// 	h.val++
// }

// func TestUnsubscribeMethod(t *testing.T) {
// 	bus := NewBus()
// 	h := &handler{val: 0}
// 	event := pb.FrameType_CLOSE
// 	bus.Subscribe(event, h.Handle)
// 	bus.Publish(&pb.Frame{Type: event})
// 	assert.Nil(t, bus.Unsubscribe(event, h.Handle))
// 	assert.NotNil(t, bus.Unsubscribe(event, h.Handle))
// 	bus.Publish(&pb.Frame{Type: event})
// 	bus.WaitAsync()
// 	assert.Equal(t, 1, h.val)
// }

// var uniqueBus = NewBus()
// var once = sync.Once{}

// func doOnce() {
// 	once.Do(func() {
// 		uniqueBus.SubscribeAsync(pb.FrameType_CLOSE, func(f *pb.Frame) {
// 		}, false)
// 	})
// }

// func BenchmarkSubscribeAsync(b *testing.B) {
// 	doOnce()
// 	for i := 0; i < 1000; i++ {
// 		uniqueBus.Publish(&pb.Frame{Type: pb.FrameType_CLOSE})
// 	}
// 	uniqueBus.WaitAsync()
// }
