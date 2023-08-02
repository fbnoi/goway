package internal

type KV[T any] struct {
	Key   string
	Value T
}
