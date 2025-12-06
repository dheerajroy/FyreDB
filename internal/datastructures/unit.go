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

func (u *Unit) GetValue() any {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.lastAccess = time.Now().UnixMilli()
	return u.value
}

func (u *Unit) SetValue(value any) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.value = value
	u.lastAccess = time.Now().UnixMilli()
}
