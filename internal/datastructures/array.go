package datastructures

type Array struct {
	BaseStructure
	data []*Unit
}

func NewArray() *Array {
	return &Array{
		data: make([]*Unit, 0, 100),
	}
}

func (array *Array) Get(index int) (*Unit, bool) {
	array.mutex.RLock()
	defer array.mutex.RUnlock()
	if index < 0 || index >= len(array.data) {
		return nil, false
	}
	return array.data[index], true
}

func (array *Array) Set(index int, unit *Unit) bool {
	array.mutex.Lock()
	defer array.mutex.Unlock()
	if index < 0 {
		return false
	}
	if index < len(array.data) {
		array.data[index] = unit
		return true
	}

	needed := index - len(array.data) + 1
	for i := 0; i < needed; i++ {
		array.data = append(array.data, nil)
	}
	array.data[index] = unit
	return true
}

func (array *Array) Append(unit *Unit) {
	array.mutex.Lock()
	defer array.mutex.Unlock()
	array.data = append(array.data, unit)
}

func (array *Array) Pop() (*Unit, bool) {
	array.mutex.Lock()
	defer array.mutex.Unlock()
	n := len(array.data)
	if n == 0 {
		return nil, false
	}
	unit := array.data[n-1]
	array.data = array.data[:n-1]
	return unit, true
}

func (array *Array) Insert(index int, unit *Unit) bool {
	array.mutex.Lock()
	defer array.mutex.Unlock()
	if index < 0 || index > len(array.data) {
		return false
	}
	array.data = append(array.data, nil)
	copy(array.data[index+1:], array.data[index:])
	array.data[index] = unit
	return true
}

func (array *Array) Delete(index int) bool {
	array.mutex.Lock()
	defer array.mutex.Unlock()
	if index < 0 || index >= len(array.data) {
		return false
	}
	array.data = append(array.data[:index], array.data[index+1:]...)
	return true
}

func (array *Array) Size() int {
	array.mutex.RLock()
	defer array.mutex.RUnlock()
	return len(array.data)
}

func (array *Array) IsEmpty() bool {
	return array.Size() == 0
}

func (array *Array) Clear() {
	array.mutex.Lock()
	defer array.mutex.Unlock()
	array.data = make([]*Unit, 0, 100)
}
