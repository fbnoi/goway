package internal

import (
	"flynoob/goway/protobuf"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestNew(t *testing.T) {
	bus := NewBus()
	if bus == nil {
		t.Log("New EventBus not created!")
		t.Fail()
	}
}

func TestHasCallback(t *testing.T) {
	bus := NewBus()
	bus.Subscribe(&protobuf.Heartbeat{}, func(proto.Message) {})
	assert.Equal(t, true, bus.HasCallback(&protobuf.Heartbeat{}))
	assert.Equal(t, false, bus.HasCallback(&protobuf.Close{}))
}

func TestSubscribe(t *testing.T) {
	bus := NewBus()
	bus.Subscribe(&protobuf.Heartbeat{}, func(proto.Message) {})
	assert.Equal(t, true, bus.HasCallback(&protobuf.Heartbeat{}))
}

func TestSubscribeOnce(t *testing.T) {
	bus := NewBus()
	bus.SubscribeOnce(&protobuf.Heartbeat{}, func(proto.Message) {})
	assert.Equal(t, true, bus.HasCallback(&protobuf.Heartbeat{}))
}

func TestSubscribeOnceAndManySubscribe(t *testing.T) {
	bus := NewBus()
	event := &protobuf.Heartbeat{}
	flag := 0
	fn := func(proto.Message) {
		flag += 1
	}
	bus.SubscribeOnce(event, fn)
	bus.Subscribe(event, fn)
	bus.Subscribe(event, fn)
	bus.Publish(event)
	assert.Equal(t, 3, flag)
}

func TestUnsubscribe(t *testing.T) {
	bus := NewBus()
	handler := func(proto.Message) {}
	bus.Subscribe(&protobuf.Heartbeat{}, handler)
	assert.Nil(t, bus.Unsubscribe(&protobuf.Heartbeat{}, handler))
	assert.NotNil(t, bus.Unsubscribe(&protobuf.Heartbeat{}, handler))
}

type handler struct {
	val int
}

func (h *handler) Handle(proto.Message) {
	h.val++
}

func TestUnsubscribeMethod(t *testing.T) {
	bus := NewBus()
	h := &handler{val: 0}
	event := &protobuf.Heartbeat{}
	bus.Subscribe(event, h.Handle)
	bus.Publish(event)
	assert.Nil(t, bus.Unsubscribe(event, h.Handle))
	assert.NotNil(t, bus.Unsubscribe(event, h.Handle))
	bus.Publish(event)
	bus.WaitAsync()
	assert.Equal(t, 1, h.val)
}

var uniqueBus = NewBus()
var once = sync.Once{}

func doOnce() {
	once.Do(func() {
		uniqueBus.SubscribeAsync(&protobuf.Heartbeat{}, func(proto.Message) {
		}, false)
	})
}

func BenchmarkSubscribeAsync(b *testing.B) {
	doOnce()
	for i := 0; i < 1000; i++ {
		uniqueBus.Publish(&protobuf.Heartbeat{})
	}
	uniqueBus.WaitAsync()
}
