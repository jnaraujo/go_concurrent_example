package mapmutex

import "sync"

type MapMutex[K comparable, V any] struct {
	mu sync.Mutex
	m  map[K]V
}

func New[K comparable, V any]() *MapMutex[K, V] {
	return &MapMutex[K, V]{
		m: make(map[K]V),
	}
}

func (mm *MapMutex[K, V]) Read(k K) (v V) {
	mm.mu.Lock()
	defer mm.mu.Unlock()
	return mm.m[k]
}

func (mm *MapMutex[K, V]) Write(k K, v V) {
	mm.mu.Lock()
	mm.m[k] = v
	mm.mu.Unlock()
}

func (mm *MapMutex[K, V]) Delete(k K) {
	mm.mu.Lock()
	delete(mm.m, k)
	mm.mu.Unlock()
}
