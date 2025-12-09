package datastructures

import "sync"

type BaseStructure struct {
	exp   int64
	mutex sync.RWMutex
}

func (datastructure *BaseStructure) SetExpiration(exp int64) {
	datastructure.mutex.Lock()
	defer datastructure.mutex.Unlock()
	datastructure.exp = exp
}

func (datastructure *BaseStructure) GetExpiration() int64 {
	datastructure.mutex.RLock()
	defer datastructure.mutex.RUnlock()
	return datastructure.exp
}
