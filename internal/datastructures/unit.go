package datastructures

import (
	"sync"
	"time"
)

type Unit struct {
	value      any
	lastAccess int64
	mutex      sync.Mutex
}

func NewUnit(value any) *Unit {
	return &Unit{
		value:      value,
		lastAccess: time.Now().UnixMilli(),
	}
}

func (unit *Unit) GetValue() any {
	unit.mutex.Lock()
	defer unit.mutex.Unlock()
	unit.lastAccess = time.Now().UnixMilli()
	return unit.value
}

func (unit *Unit) SetValue(value any) {
	unit.mutex.Lock()
	defer unit.mutex.Unlock()
	unit.value = value
	unit.lastAccess = time.Now().UnixMilli()
}
