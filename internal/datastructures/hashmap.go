package datastructures

import "sync"

type HashMap struct {
	data map[string]any
	mu   sync.RWMutex
}

func NewHashMap() *HashMap {
	return &HashMap{
		data: make(map[string]any),
	}
}

func (hm *HashMap) Set(key string, value any) {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	hm.data[key] = value
}

func (hm *HashMap) Get(key string) (any, bool) {
	hm.mu.RLock()
	defer hm.mu.RUnlock()
	value, exists := hm.data[key]
	return value, exists
}

func (hm *HashMap) Delete(key string) {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	delete(hm.data, key)
}

func (hm *HashMap) Keys() []string {
	hm.mu.RLock()
	defer hm.mu.RUnlock()
	keys := make([]string, 0, len(hm.data))
	for key := range hm.data {
		keys = append(keys, key)
	}
	return keys
}

func (hm *HashMap) Size() int {
	hm.mu.RLock()
	defer hm.mu.RUnlock()
	return len(hm.data)
}

func (hm *HashMap) Clear() {
	hm.mu.Lock()
	defer hm.mu.Unlock()
	hm.data = make(map[string]any)
}
