package goway

type KV[T any] struct {
	Key   string
	Value T
}
