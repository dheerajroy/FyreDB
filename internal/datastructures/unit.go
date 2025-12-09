package datastructures

import (
	"sync"
	"time"
)

type Unit struct {
	value      any
	lastAccess int64
	createdOn  int64
	mutex      sync.RWMutex
}

func NewUnit(value any) *Unit {
	return &Unit{
		value:      value,
		lastAccess: time.Now().UnixMilli(),
		createdOn:  time.Now().UnixMilli(),
	}
}

func (unit *Unit) GetValue() any {
	unit.mutex.RLock()
	defer unit.mutex.RUnlock()
	unit.lastAccess = time.Now().UnixMilli()
	return unit.value
}

func (unit *Unit) SetValue(value any) {
	unit.mutex.Lock()
	defer unit.mutex.Unlock()
	unit.value = value
	unit.lastAccess = time.Now().UnixMilli()
}

func (unit *Unit) GetLastAccess() int64 {
	unit.mutex.RLock()
	defer unit.mutex.RUnlock()
	return unit.lastAccess
}

func (unit *Unit) GetCreatedOn() int64 {
	unit.mutex.RLock()
	defer unit.mutex.RUnlock()
	return unit.createdOn
}
