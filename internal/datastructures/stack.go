package datastructures

type Stack struct {
	BaseStructure
	data []*Unit
}

func NewStack() *Stack {
	return &Stack{
		data: make([]*Unit, 0, 100),
	}
}

func (stack *Stack) Push(unit *Unit) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.data = append(stack.data, unit)
}

func (stack *Stack) Pop() (*Unit, bool) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	dataLength := len(stack.data)
	if dataLength == 0 {
		return nil, false
	}
	unit := stack.data[dataLength-1]
	stack.data = stack.data[:dataLength-1]
	return unit, true
}

func (stack *Stack) Peek() (*Unit, bool) {
	stack.mutex.RLock()
	defer stack.mutex.RUnlock()
	if len(stack.data) == 0 {
		return nil, false
	}
	return stack.data[len(stack.data)-1], true
}

func (stack *Stack) Size() int {
	stack.mutex.RLock()
	defer stack.mutex.RUnlock()
	return len(stack.data)
}

func (stack *Stack) IsEmpty() bool {
	stack.mutex.RLock()
	defer stack.mutex.RUnlock()
	return len(stack.data) == 0
}

func (stack *Stack) Clear() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.data = make([]*Unit, 0, 100)
}
