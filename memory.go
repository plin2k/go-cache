package go_cache

import (
	"sync"
	"time"
)

type dataCache struct {
	expiration time.Time
	data       interface{}
}

type memoryCache struct {
	sync.RWMutex
	data map[string]dataCache
}

func NewMemory() (*memoryCache, error) {
	return &memoryCache{
		data: map[string]dataCache{},
	}, nil
}

func (m *memoryCache) Get(key string) (interface{}, error) {
	m.RLock()
	val, exists := m.data[key]
	m.RUnlock()

	if exists && time.Now().After(val.expiration) {
		m.Lock()
		delete(m.data, key)
		m.Unlock()
		exists = false
	}

	if !exists {
		return nil, errNotFound
	}

	return val.data, nil
}

func (m *memoryCache) Set(key string, val interface{}, duration time.Duration) error {
	var nowTime = time.Now()

	if duration == 0 {
		nowTime = nowTime.Add(1 << 16 * time.Hour)
	}

	m.Lock()
	m.data[key] = dataCache{
		data:       val,
		expiration: nowTime.Add(duration),
	}
	m.Unlock()

	return nil
}

func (m *memoryCache) Delete(key string) error {
	m.Lock()
	delete(m.data, key)
	m.Unlock()

	return nil
}

func (m *memoryCache) Flush() error {
	m.Lock()
	m.data = map[string]dataCache{}
	m.Unlock()

	return nil
}
