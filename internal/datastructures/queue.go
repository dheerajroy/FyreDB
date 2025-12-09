package datastructures

type Queue struct {
	BaseStructure
	data []*Unit
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]*Unit, 0, 100),
	}
}

func (queue *Queue) Enqueue(unit *Unit) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	queue.data = append(queue.data, unit)
}

func (queue *Queue) Dequeue() (*Unit, bool) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	if len(queue.data) == 0 {
		return nil, false
	}
	unit := queue.data[0]
	queue.data = queue.data[1:]
	return unit, true
}

func (queue *Queue) Peek() (*Unit, bool) {
	queue.mutex.RLock()
	defer queue.mutex.RUnlock()
	if len(queue.data) == 0 {
		return nil, false
	}
	return queue.data[0], true
}

func (queue *Queue) Size() int {
	queue.mutex.RLock()
	defer queue.mutex.RUnlock()
	return len(queue.data)
}

func (queue *Queue) IsEmpty() bool {
	queue.mutex.RLock()
	defer queue.mutex.RUnlock()
	return len(queue.data) == 0
}

func (queue *Queue) Clear() {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	queue.data = make([]*Unit, 0, 100)
}
