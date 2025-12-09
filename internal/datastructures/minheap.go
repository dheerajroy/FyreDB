package datastructures

type Item struct {
	Value *Unit
	Score float64
}

type MinHeap struct {
	BaseStructure
	data []*Item
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		data: make([]*Item, 0, 100),
	}
}

func (minHeap *MinHeap) Size() int {
	minHeap.mutex.RLock()
	defer minHeap.mutex.RUnlock()
	return len(minHeap.data)
}

func (minHeap *MinHeap) Push(item *Item) {
	minHeap.mutex.Lock()
	defer minHeap.mutex.Unlock()
	minHeap.data = append(minHeap.data, item)
	minHeap.up(len(minHeap.data) - 1)
}

func (minHeap *MinHeap) Pop() *Item {
	minHeap.mutex.Lock()
	defer minHeap.mutex.Unlock()
	n := len(minHeap.data) - 1
	if n < 0 {
		return nil
	}

	minHeap.data[0], minHeap.data[n] = minHeap.data[n], minHeap.data[0]

	item := minHeap.data[n]
	minHeap.data = minHeap.data[:n]

	minHeap.down(0, len(minHeap.data))

	return item
}

func (minHeap *MinHeap) Peek() *Item {
	minHeap.mutex.RLock()
	defer minHeap.mutex.RUnlock()
	if len(minHeap.data) == 0 {
		return nil
	}
	return minHeap.data[0]
}

func (minHeap *MinHeap) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || minHeap.data[i].Score <= minHeap.data[j].Score {
			break
		}
		minHeap.data[i], minHeap.data[j] = minHeap.data[j], minHeap.data[i]
		j = i
	}
}

func (minHeap *MinHeap) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n {
			break
		}
		j := j1

		if j2 := j1 + 1; j2 < n && minHeap.data[j2].Score < minHeap.data[j1].Score {
			j = j2
		}

		if minHeap.data[j].Score >= minHeap.data[i].Score {
			break
		}

		minHeap.data[i], minHeap.data[j] = minHeap.data[j], minHeap.data[i]
		i = j
	}
	return i > i0
}

func (minHeap *MinHeap) Clear() {
	minHeap.mutex.Lock()
	defer minHeap.mutex.Unlock()
	minHeap.data = make([]*Item, 0, 100)
}
