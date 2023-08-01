package goway

import (
	"reflect"
	"sync"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

type HandleFunc func(proto.Message)

type BusSubscriber interface {
	Subscribe(typ proto.Message, fn HandleFunc)
	SubscribeAsync(typ proto.Message, fn HandleFunc, transactional bool)
	SubscribeOnce(typ proto.Message, fn HandleFunc)
	SubscribeOnceAsync(typ proto.Message, fn HandleFunc)
	Unsubscribe(typ proto.Message, fn HandleFunc) error
}

type BusPublisher interface {
	Publish(proto.Message)
}

type BusController interface {
	HasCallback(typ proto.Message) bool
	WaitAsync()
}

type Bus interface {
	BusController
	BusSubscriber
	BusPublisher
}

func NewBus() Bus {
	b := &EventBus{
		make(map[string][]*eventHandler),
		sync.RWMutex{},
		sync.WaitGroup{},
	}
	return b
}

type EventBus struct {
	handlers map[string][]*eventHandler
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

func (bus *EventBus) doSubscribe(typ proto.Message, handler *eventHandler) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	bus.handlers[typ] = append(bus.handlers[typ], handler)
}

func (bus *EventBus) Subscribe(typ proto.Message, fn HandleFunc) {
	bus.doSubscribe(typ, &eventHandler{fn, false, false, false, sync.Mutex{}})
}

func (bus *EventBus) SubscribeAsync(typ proto.Message, fn HandleFunc, transactional bool) {
	bus.doSubscribe(typ, &eventHandler{fn, false, true, transactional, sync.Mutex{}})
}

func (bus *EventBus) SubscribeOnce(typ proto.Message, fn HandleFunc) {
	bus.doSubscribe(typ, &eventHandler{fn, true, false, false, sync.Mutex{}})
}

func (bus *EventBus) SubscribeOnceAsync(typ proto.Message, fn HandleFunc) {
	bus.doSubscribe(typ, &eventHandler{fn, true, true, false, sync.Mutex{}})
}

func (bus *EventBus) Unsubscribe(typ proto.Message, fn HandleFunc) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()
	if _, ok := bus.handlers[typ]; ok && len(bus.handlers[typ]) > 0 {
		bus.removeHandler(typ, bus.findHandlerIdx(typ, fn))
		return nil
	}
	return errors.Errorf("frame %v doesn't exist", typ)
}

func (bus *EventBus) Publish(frame proto.Message) {
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

func (bus *EventBus) HasCallback(typ proto.Message) bool {
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

func (bus *EventBus) doPublish(handler *eventHandler, frame proto.Message) {
	handler.handleFunc(frame)
}

func (bus *EventBus) doPublishAsync(handler *eventHandler, frame proto.Message) {
	defer bus.wg.Done()
	if handler.transactional {
		defer handler.Unlock()
	}
	bus.doPublish(handler, frame)
}

func (bus *EventBus) removeHandler(typ proto.Message, i int) {
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

func (bus *EventBus) findHandlerIdx(typ proto.Message, fn HandleFunc) int {
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
