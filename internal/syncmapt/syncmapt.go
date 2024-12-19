package syncmapt

import (
	"sync"
)

type SyncMap[K comparable, V any] struct {
	innerMap sync.Map
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.innerMap.Delete(key)
}

func (m *SyncMap[K, V]) Load(key K) (*V, bool) {
	var value *V = nil

	v, ok := m.innerMap.Load(key)

	if !ok {
		return value, ok
	}

	value, ok = v.(*V)

	return value, ok
}

func (m *SyncMap[K, V]) Store(key K, value *V) {
	m.innerMap.Store(key, value)
}

func (m *SyncMap[K, V]) CompareAndSwap(key K, oldv *V, newv *V) bool {
	return m.innerMap.CompareAndSwap(key, oldv, newv)
}
