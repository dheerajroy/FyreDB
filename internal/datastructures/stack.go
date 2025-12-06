package datastructures

import "sync"

type Stack struct {
	data []*Unit
	mutex   sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{
		data: make([]*Unit, 0, 100),
	}
}

func (s *Stack) Push(value any) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = append(s.data, NewUnit(value))
}

func (s *Stack) Pop() (any, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	dataLength := len(s.data)
	if dataLength == 0 {
		return nil, false
	}
	unit := s.data[dataLength-1]
	s.data = s.data[:dataLength-1]
	return unit.GetValue(), true
}

func (s *Stack) Peek() (any, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if len(s.data) == 0 {
		return nil, false
	}
	return s.data[len(s.data)-1].GetValue(), true
}

func (s *Stack) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.data) == 0
}

func (s *Stack) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data = make([]*Unit, 0, 100)
}
