package datastructures

type HashMap struct {
	BaseStructure
	data map[string]*Unit
}

func NewHashMap() *HashMap {
	return &HashMap{
		data: make(map[string]*Unit),
	}
}

func (hashmap *HashMap) Set(key string, value *Unit) {
	hashmap.mutex.Lock()
	defer hashmap.mutex.Unlock()
	hashmap.data[key] = value
}

func (hashmap *HashMap) Get(key string) (*Unit, bool) {
	hashmap.mutex.RLock()
	defer hashmap.mutex.RUnlock()
	value, exists := hashmap.data[key]
	return value, exists
}

func (hashmap *HashMap) Delete(key string) {
	hashmap.mutex.Lock()
	defer hashmap.mutex.Unlock()
	delete(hashmap.data, key)
}

func (hashmap *HashMap) Keys() []string {
	hashmap.mutex.RLock()
	defer hashmap.mutex.RUnlock()
	keys := make([]string, 0, len(hashmap.data))
	for key := range hashmap.data {
		keys = append(keys, key)
	}
	return keys
}

func (hashmap *HashMap) Size() int {
	hashmap.mutex.RLock()
	defer hashmap.mutex.RUnlock()
	return len(hashmap.data)
}

func (hashmap *HashMap) Clear() {
	hashmap.mutex.Lock()
	defer hashmap.mutex.Unlock()
	hashmap.data = make(map[string]*Unit)
}
