package protobuf

import (
	"sync"

	"google.golang.org/protobuf/proto"
)

func NewFactory[T proto.Message]() *factory[T] {
	return &factory[T]{
		pool: sync.Pool{
			New: func() any {
				return new(T)
			},
		},
	}
}

type factory[T proto.Message] struct {
	pool sync.Pool
}

func (f *factory[T]) Get() T {
	return f.pool.Get().(T)
}

func (f *factory[T]) Put(m T) {
	f.pool.Put(m)
}

func (f *factory[T]) FromBytes(bs []byte) (m T, err error) {
	m = f.Get()
	err = proto.Unmarshal(bs, m)

	return
}
