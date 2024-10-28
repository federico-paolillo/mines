package syncmapt

import "sync"

type SyncMap[K comparable, V any] struct {
	innerMap sync.Map
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.innerMap.Delete(key)
}

func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.innerMap.Load(key)

	if !ok {
		return value, ok
	}

	return v.(V), ok
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.innerMap.Store(key, value)
}

func (m *SyncMap[K, V]) CompareAndSwap(key K, old V, new V) bool {
	return m.innerMap.CompareAndSwap(key, old, new)
}
