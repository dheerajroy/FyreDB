package datastructures

import "sync"

type Item struct {
	Value *Unit
	Score float64
}

type MinHeap struct {
	data  []*Item
	mutex sync.RWMutex
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: make([]*Item, 0, 100),
	}
}

func (heap *MinHeap) Len() int {
	heap.mutex.RLock()
	defer heap.mutex.RUnlock()
	return len(heap.data)
}

func (heap *MinHeap) Push(item *Item) {
	heap.mutex.Lock()
	defer heap.mutex.Unlock()
	heap.data = append(heap.data, item)
	heap.up(len(heap.data) - 1)
}

func (heap *MinHeap) Pop() *Item {
	heap.mutex.Lock()
	defer heap.mutex.Unlock()
	n := len(heap.data) - 1
	if n < 0 {
		return nil
	}

	heap.data[0], heap.data[n] = heap.data[n], heap.data[0]

	item := heap.data[n]
	heap.data = heap.data[:n]

	heap.down(0, len(heap.data))

	return item
}

func (heap *MinHeap) Peek() *Item {
	heap.mutex.RLock()
	defer heap.mutex.RUnlock()
	if len(heap.data) == 0 {
		return nil
	}
	return heap.data[0]
}

func (heap *MinHeap) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || heap.data[i].Score <= heap.data[j].Score {
			break
		}
		heap.data[i], heap.data[j] = heap.data[j], heap.data[i]
		j = i
	}
}

func (heap *MinHeap) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n {
			break
		}
		j := j1

		if j2 := j1 + 1; j2 < n && heap.data[j2].Score < heap.data[j1].Score {
			j = j2
		}

		if heap.data[j].Score >= heap.data[i].Score {
			break
		}

		heap.data[i], heap.data[j] = heap.data[j], heap.data[i]
		i = j
	}
	return i > i0
}

func (heap *MinHeap) Clear() {
	heap.mutex.Lock()
	defer heap.mutex.Unlock()
	heap.data = make([]*Item, 0, 100)
}
