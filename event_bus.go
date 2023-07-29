package goway

import (
	pb "flynoob/goway/protobuf"
	"reflect"
	"sync"

	"github.com/pkg/errors"
)

type HandleFunc func(*pb.Frame)

type BusSubscriber interface {
	Subscribe(typ pb.FrameType, fn HandleFunc)
	SubscribeAsync(typ pb.FrameType, fn HandleFunc, transactional bool)
	SubscribeOnce(typ pb.FrameType, fn HandleFunc)
	SubscribeOnceAsync(typ pb.FrameType, fn HandleFunc)
	Unsubscribe(typ pb.FrameType, fn HandleFunc) error
}

type BusPublisher interface {
	Publish(*pb.Frame)
}

type BusController interface {
	HasCallback(typ pb.FrameType) bool
	WaitAsync()
}

type Bus interface {
	BusController
	BusSubscriber
	BusPublisher
}

func NewBus() Bus {
	b := &EventBus{
		make(map[pb.FrameType][]*eventHandler),
		sync.RWMutex{},
		sync.WaitGroup{},
	}
	return b
}

type EventBus struct {
	handlers map[pb.FrameType][]*eventHandler
	lock     sync.RWMutex
	wg       sync.WaitGroup
}

type eventHandler struct {
	handleFunc    HandleFunc
	flagOnce      bool
	async         bool
	transactional bool
	sync.Mutex
}

func (bus *EventBus) doSubscribe(typ pb.FrameType, handler *eventHandler) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	bus.handlers[typ] = append(bus.handlers[typ], handler)
}

func (bus *EventBus) Subscribe(typ pb.FrameType, fn HandleFunc) {
	bus.doSubscribe(typ, &eventHandler{fn, false, false, false, sync.Mutex{}})
}

func (bus *EventBus) SubscribeAsync(typ pb.FrameType, fn HandleFunc, transactional bool) {
	bus.doSubscribe(typ, &eventHandler{fn, false, true, transactional, sync.Mutex{}})
}

func (bus *EventBus) SubscribeOnce(typ pb.FrameType, fn HandleFunc) {
	bus.doSubscribe(typ, &eventHandler{fn, true, false, false, sync.Mutex{}})
}

func (bus *EventBus) SubscribeOnceAsync(typ pb.FrameType, fn HandleFunc) {
	bus.doSubscribe(typ, &eventHandler{fn, true, true, false, sync.Mutex{}})
}

func (bus *EventBus) Unsubscribe(typ pb.FrameType, fn HandleFunc) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()
	if _, ok := bus.handlers[typ]; ok && len(bus.handlers[typ]) > 0 {
		bus.removeHandler(typ, bus.findHandlerIdx(typ, fn))
		return nil
	}
	return errors.Errorf("frame %v doesn't exist", typ)
}

func (bus *EventBus) Publish(frame *pb.Frame) {
	bus.lock.RLock()
	defer bus.lock.RUnlock()
	if handlers, ok := bus.handlers[frame.Type]; ok && 0 < len(handlers) {
		onceIdx := []int{}
		for i, handler := range handlers {
			if handler.flagOnce {
				onceIdx = append(onceIdx, i)
			}
			if !handler.async {
				bus.doPublish(handler, frame)
			} else {
				bus.wg.Add(1)
				if handler.transactional {
					bus.lock.Unlock()
					handler.Lock()
					bus.lock.Lock()
				}
				go bus.doPublishAsync(handler, frame)
			}
		}
		for _, i := range onceIdx {
			bus.removeHandler(frame.Type, i)
		}
	}
}

func (bus *EventBus) HasCallback(typ pb.FrameType) bool {
	bus.lock.Lock()
	defer bus.lock.Unlock()
	_, ok := bus.handlers[typ]
	if ok {
		return len(bus.handlers[typ]) > 0
	}
	return false
}

func (bus *EventBus) WaitAsync() {
	bus.wg.Wait()
}

func (bus *EventBus) doPublish(handler *eventHandler, frame *pb.Frame) {
	handler.handleFunc(frame)
}

func (bus *EventBus) doPublishAsync(handler *eventHandler, frame *pb.Frame) {
	defer bus.wg.Done()
	if handler.transactional {
		defer handler.Unlock()
	}
	bus.doPublish(handler, frame)
}

func (bus *EventBus) removeHandler(typ pb.FrameType, i int) {
	if _, ok := bus.handlers[typ]; !ok {
		return
	}
	l := len(bus.handlers[typ])

	if !(0 <= i && i < l) {
		return
	}

	copy(bus.handlers[typ][i:], bus.handlers[typ][i+1:])
	bus.handlers[typ][l-1] = nil
	bus.handlers[typ] = bus.handlers[typ][:l-1]
}

func (bus *EventBus) findHandlerIdx(typ pb.FrameType, fn HandleFunc) int {
	if _, ok := bus.handlers[typ]; ok {
		for i, v := range bus.handlers[typ] {
			sf1 := reflect.ValueOf(v.handleFunc)
			sf2 := reflect.ValueOf(fn)
			if sf1.Type() == sf2.Type() && sf1.Pointer() == sf2.Pointer() {
				return i
			}
		}
	}
	return -1
}
