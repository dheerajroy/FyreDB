package datastructures

import "sync"

type Queue struct {
	data []*Unit
	mutex   sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]*Unit, 0, 100),
	}
}

func (q *Queue) Enqueue(value any) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.data = append(q.data, NewUnit(value))
}

func (q *Queue) Dequeue() (any, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.data) == 0 {
		return nil, false
	}
	value := q.data[0]
	q.data = q.data[1:]
	return value, true
}

func (q *Queue) Peek() (any, bool) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	if len(q.data) == 0 {
		return nil, false
	}
	return q.data[0], true
}

func (q *Queue) Size() int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return len(q.data) == 0
}

func (q *Queue) Clear() {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.data = make([]*Unit, 0, 100)
}
