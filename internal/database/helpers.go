package database

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dheerajroy/fyredb/internal/datastructures"
)

func parsePrimitive(s string) any {
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i
	}
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}
	if s == "true" {
		return true
	}
	if s == "false" {
		return false
	}
	return s
}

func useNode(current CurrentNode, name string) bool {
	switch current.Type {
	case HashMapType:
		hm := current.Value.(*datastructures.HashMap)
		unit, ok := hm.Get(name)
		if !ok {
			fmt.Println("No such key in hashmap:", name)
			return false
		}

		switch v := unit.GetValue().(type) {
		case *datastructures.HashMap:
			nodeStack = append(nodeStack, CurrentNode{Type: HashMapType, Value: v})
			nameStack = append(nameStack, name)
			fmt.Println("Now using hashmap:", name)
			return true
		case *datastructures.Array:
			nodeStack = append(nodeStack, CurrentNode{Type: ArrayType, Value: v})
			nameStack = append(nameStack, name)
			fmt.Println("Now using array:", name)
			return true
		case *datastructures.Stack:
			nodeStack = append(nodeStack, CurrentNode{Type: StackType, Value: v})
			nameStack = append(nameStack, name)
			fmt.Println("Now using stack:", name)
			return true
		case *datastructures.Queue:
			nodeStack = append(nodeStack, CurrentNode{Type: QueueType, Value: v})
			nameStack = append(nameStack, name)
			fmt.Println("Now using queue:", name)
			return true
		case *datastructures.MinHeap:
			nodeStack = append(nodeStack, CurrentNode{Type: MinHeapType, Value: v})
			nameStack = append(nameStack, name)
			fmt.Println("Now using minheap:", name)
			return true
		default:
			fmt.Println("Value at key is not a container you can use")
			return false
		}

	case ArrayType:
		arr := current.Value.(*datastructures.Array)
		idx, err := strconv.Atoi(name)
		if err != nil {
			fmt.Println("Invalid index:", name)
			return false
		}
		unit, ok := arr.Get(idx)
		if !ok {
			fmt.Println("Index out of range")
			return false
		}
		switch v := unit.GetValue().(type) {
		case *datastructures.HashMap:
			nodeStack = append(nodeStack, CurrentNode{Type: HashMapType, Value: v})
			nameStack = append(nameStack, name)
			fmt.Println("Now using hashmap at index", idx)
			return true
		default:
			fmt.Println("Value at index is not a container you can use")
			return false
		}

	default:
		fmt.Println("use is only supported for hashmap and array container types")
		return false
	}
}

func getNode(current CurrentNode, name string) {
	switch current.Type {
	case HashMapType:
		hm := current.Value.(*datastructures.HashMap)
		unit, ok := hm.Get(name)
		if !ok {
			fmt.Println("Key not found:", name)
			return
		}
		fmt.Printf("%s -> %T : %v\n", name, unit.GetValue(), unit.GetValue())

	case ArrayType:
		arr := current.Value.(*datastructures.Array)
		idx, err := strconv.Atoi(name)
		if err != nil {
			fmt.Println("Invalid index:", name)
			return
		}
		unit, ok := arr.Get(idx)
		if !ok {
			fmt.Println("Index out of range")
			return
		}
		fmt.Printf("[%d] -> %T : %v\n", idx, unit.GetValue(), unit.GetValue())

	case StackType:
		st := current.Value.(*datastructures.Stack)
		if v, ok := st.Peek(); ok {
			fmt.Printf("top -> %T : %v\n", v.GetValue(), v.GetValue())
		} else {
			fmt.Println("stack empty")
		}

	case QueueType:
		q := current.Value.(*datastructures.Queue)
		if v, ok := q.Peek(); ok {
			fmt.Printf("front -> %T : %v\n", v.GetValue(), v.GetValue())
		} else {
			fmt.Println("queue empty")
		}

	case MinHeapType:
		mh := current.Value.(*datastructures.MinHeap)
		if v := mh.Peek(); v != nil {
			fmt.Printf("peek -> %T : %v (score=%f)\n", v.Value.GetValue(), v.Value.GetValue(), v.Score)
		} else {
			fmt.Println("minheap empty")
		}

	default:
		fmt.Println("get not supported for this node type")
	}
}

func delNode(current CurrentNode, name string) {
	switch current.Type {
	case HashMapType:
		hm := current.Value.(*datastructures.HashMap)
		hm.Delete(name)
		fmt.Println("Deleted key:", name)

	case ArrayType:
		arr := current.Value.(*datastructures.Array)
		idx, err := strconv.Atoi(name)
		if err != nil {
			fmt.Println("Invalid index:", name)
			return
		}
		if arr.Delete(idx) {
			fmt.Println("Deleted index:", idx)
		} else {
			fmt.Println("Index out of range")
		}

	case StackType:
		if name == "0" || strings.ToLower(name) == "top" {
			if v, ok := current.Value.(*datastructures.Stack).Pop(); ok {
				fmt.Printf("Popped -> %T : %v\n", v.GetValue(), v.GetValue())
			} else {
				fmt.Println("stack empty")
			}
		} else {
			fmt.Println("stack only supports deleting the top element (use 'del top' or 'del 0')")
		}

	case QueueType:
		if name == "0" || strings.ToLower(name) == "front" {
			if v, ok := current.Value.(*datastructures.Queue).Dequeue(); ok {
				fmt.Printf("Dequeued -> %T : %v\n", v.GetValue(), v.GetValue())
			} else {
				fmt.Println("queue empty")
			}
		} else {
			fmt.Println("queue only supports deleting the front element (use 'del front' or 'del 0')")
		}

	case MinHeapType:
		if name == "0" || strings.ToLower(name) == "top" || name == "pop" {
			if v := current.Value.(*datastructures.MinHeap).Pop(); v != nil {
				fmt.Printf("Popped -> %T : %v (score=%f)\n", v.Value.GetValue(), v.Value.GetValue(), v.Score)
			} else {
				fmt.Println("minheap empty")
			}
		} else {
			fmt.Println("minheap only supports popping the top element (use 'del top')")
		}

	default:
		fmt.Println("del not supported for this node type")
	}
}

func listNode(current CurrentNode) {
	switch current.Type {
	case HashMapType:
		hm := current.Value.(*datastructures.HashMap)
		keys := hm.Keys()
		if len(keys) == 0 {
			fmt.Println("(empty hashmap)")
			return
		}
		for _, k := range keys {
			if u, ok := hm.Get(k); ok {
				fmt.Printf("%s -> %T\n", k, u.GetValue())
			}
		}

	case ArrayType:
		arr := current.Value.(*datastructures.Array)
		sz := arr.Size()
		if sz == 0 {
			fmt.Println("(empty array)")
			return
		}
		for i := 0; i < sz; i++ {
			if u, ok := arr.Get(i); ok {
				fmt.Printf("[%d] -> %T\n", i, u.GetValue())
			}
		}

	case StackType:
		st := current.Value.(*datastructures.Stack)
		fmt.Printf("stack size: %d\n", st.Size())

	case QueueType:
		q := current.Value.(*datastructures.Queue)
		fmt.Printf("queue size: %d\n", q.Size())

	case MinHeapType:
		mh := current.Value.(*datastructures.MinHeap)
		fmt.Printf("minheap size: %d\n", mh.Size())

	default:
		fmt.Println("Nothing to list in this node type")
	}
}

func pushValue(current CurrentNode, raw string) {
	unit := datastructures.NewUnit(parsePrimitive(raw))
	switch current.Type {
	case ArrayType:
		current.Value.(*datastructures.Array).Append(unit)
		fmt.Println("Appended to array")
	case StackType:
		current.Value.(*datastructures.Stack).Push(unit)
		fmt.Println("Pushed to stack")
	case QueueType:
		current.Value.(*datastructures.Queue).Enqueue(unit)
		fmt.Println("Enqueued to queue")
	case MinHeapType:
		current.Value.(*datastructures.MinHeap).Push(&datastructures.Item{Value: unit, Score: 0})
		fmt.Println("Pushed to minheap with score 0")
	default:
		fmt.Println("push supported only for array/stack/queue/minheap")
	}
}

func popNode(current CurrentNode) {
	switch current.Type {
	case StackType:
		if v, ok := current.Value.(*datastructures.Stack).Pop(); ok {
			fmt.Printf("Popped -> %T : %v\n", v.GetValue(), v.GetValue())
		} else {
			fmt.Println("stack empty")
		}
	case QueueType:
		if v, ok := current.Value.(*datastructures.Queue).Dequeue(); ok {
			fmt.Printf("Dequeued -> %T : %v\n", v.GetValue(), v.GetValue())
		} else {
			fmt.Println("queue empty")
		}
	case ArrayType:
		if v, ok := current.Value.(*datastructures.Array).Pop(); ok {
			fmt.Printf("Popped -> %T : %v\n", v.GetValue(), v.GetValue())
		} else {
			fmt.Println("array empty")
		}
	case MinHeapType:
		if v := current.Value.(*datastructures.MinHeap).Pop(); v != nil {
			fmt.Printf("Popped -> %T : %v (score=%f)\n", v.Value.GetValue(), v.Value.GetValue(), v.Score)
		} else {
			fmt.Println("minheap empty")
		}
	default:
		fmt.Println("pop/dequeue not supported for this node type")
	}
}

func clearNode(current CurrentNode) {
	switch current.Type {
	case HashMapType:
		current.Value.(*datastructures.HashMap).Clear()
		fmt.Println("cleared hashmap")
	case ArrayType:
		current.Value.(*datastructures.Array).Clear()
		fmt.Println("cleared array")
	case StackType:
		current.Value.(*datastructures.Stack).Clear()
		fmt.Println("cleared stack")
	case QueueType:
		current.Value.(*datastructures.Queue).Clear()
		fmt.Println("cleared queue")
	case MinHeapType:
		current.Value.(*datastructures.MinHeap).Clear()
		fmt.Println("cleared minheap")
	default:
		fmt.Println("clear not supported for this node type")
	}
}

func sizeNode(current CurrentNode) {
	switch current.Type {
	case HashMapType:
		fmt.Printf("size: %d\n", current.Value.(*datastructures.HashMap).Size())
	case ArrayType:
		fmt.Printf("size: %d\n", current.Value.(*datastructures.Array).Size())
	case StackType:
		fmt.Printf("size: %d\n", current.Value.(*datastructures.Stack).Size())
	case QueueType:
		fmt.Printf("size: %d\n", current.Value.(*datastructures.Queue).Size())
	case MinHeapType:
		fmt.Printf("size: %d\n", current.Value.(*datastructures.MinHeap).Size())
	default:
		fmt.Println("size not supported for this node type")
	}
}

func insertIndex(current CurrentNode, index int, raw string) {
	if current.Type != ArrayType {
		fmt.Println("insert supported only for array")
		return
	}
	unit := datastructures.NewUnit(parsePrimitive(raw))
	if current.Value.(*datastructures.Array).Insert(index, unit) {
		fmt.Println("Inserted at index", index)
	} else {
		fmt.Println("Insert failed: invalid index")
	}
}

func setIndex(current CurrentNode, index int, raw string) {
	if current.Type != ArrayType {
		fmt.Println("setindex supported only for array")
		return
	}
	unit := datastructures.NewUnit(parsePrimitive(raw))
	if current.Value.(*datastructures.Array).Set(index, unit) {
		fmt.Println("Set index", index)
	} else {
		fmt.Println("Set failed: invalid index")
	}
}

func peekNode(current CurrentNode) {
	switch current.Type {
	case StackType:
		if v, ok := current.Value.(*datastructures.Stack).Peek(); ok {
			fmt.Printf("top -> %T : %v\n", v.GetValue(), v.GetValue())
		} else {
			fmt.Println("stack empty")
		}
	case QueueType:
		if v, ok := current.Value.(*datastructures.Queue).Peek(); ok {
			fmt.Printf("front -> %T : %v\n", v.GetValue(), v.GetValue())
		} else {
			fmt.Println("queue empty")
		}
	case MinHeapType:
		if v := current.Value.(*datastructures.MinHeap).Peek(); v != nil {
			fmt.Printf("peek -> %T : %v (score=%f)\n", v.Value.GetValue(), v.Value.GetValue(), v.Score)
		} else {
			fmt.Println("minheap empty")
		}
	default:
		fmt.Println("peek not supported for this node type")
	}
}

func setExpiration(current CurrentNode, ms int64) {
	switch current.Type {
	case HashMapType:
		current.Value.(*datastructures.HashMap).SetExpiration(ms)
		fmt.Println("Set expiration")
	case ArrayType:
		current.Value.(*datastructures.Array).SetExpiration(ms)
		fmt.Println("Set expiration")
	case StackType:
		current.Value.(*datastructures.Stack).SetExpiration(ms)
		fmt.Println("Set expiration")
	case QueueType:
		current.Value.(*datastructures.Queue).SetExpiration(ms)
		fmt.Println("Set expiration")
	case MinHeapType:
		current.Value.(*datastructures.MinHeap).SetExpiration(ms)
		fmt.Println("Set expiration")
	default:
		fmt.Println("expire not supported for this node type")
	}
}

func getExpiration(current CurrentNode) {
	switch current.Type {
	case HashMapType:
		fmt.Printf("expiration: %d\n", current.Value.(*datastructures.HashMap).GetExpiration())
	case ArrayType:
		fmt.Printf("expiration: %d\n", current.Value.(*datastructures.Array).GetExpiration())
	case StackType:
		fmt.Printf("expiration: %d\n", current.Value.(*datastructures.Stack).GetExpiration())
	case QueueType:
		fmt.Printf("expiration: %d\n", current.Value.(*datastructures.Queue).GetExpiration())
	case MinHeapType:
		fmt.Printf("expiration: %d\n", current.Value.(*datastructures.MinHeap).GetExpiration())
	default:
		fmt.Println("getexp not supported for this node type")
	}
}

func pushToMinHeap(current CurrentNode, score float64, raw string) {
	if current.Type != MinHeapType {
		fmt.Println("push <score> supported only for minheap")
		return
	}
	unit := datastructures.NewUnit(parsePrimitive(raw))
	current.Value.(*datastructures.MinHeap).Push(&datastructures.Item{Value: unit, Score: score})
	fmt.Println("Pushed to minheap with score", score)
}

func printHelp() {
	fmt.Println("Commands:")
	fmt.Println("  create <name> <type>    - Create a named datastructure in current hashmap. Example: create users hashmap")
	fmt.Println("  use <key|index>         - Enter a child container (use key for hashmap, index for array). Example: use users")
	fmt.Println("  back                    - Go to parent container")
	fmt.Println("  root                    - Return to root store")
	fmt.Println("  set <key> <value>       - Set primitive value in hashmap. Example: set id 123")
	fmt.Println("  get <key|index>         - Get value (or peek for stack/queue/minheap). Example: get id")
	fmt.Println("  del <key|index|top>     - Delete key or pop element. Examples: del id, del 0, del top")
	fmt.Println("  ls | keys               - List keys or indices in the current container")
	fmt.Println("  clear                   - Clear the current container")
	fmt.Println("  size                    - Show size of current container")
	fmt.Println("  insert <i> <v>          - Insert into array at index i. Example: insert 1 hello")
	fmt.Println("  setindex <i> <v>        - Set (or extend) array index i. Example: setindex 2 world")
	fmt.Println("  push <value>            - Push/append value for stack/queue/array; for minheap use: push <score> <value>. Example: push 42 OR push 1.5 item")
	fmt.Println("  pop | dequeue           - Pop from stack/array or dequeue from queue; pop minheap to remove top")
	fmt.Println("  peek                    - Show top/front/peek without removing")
	fmt.Println("  expire <ms>             - Set expiration (milliseconds) on container")
	fmt.Println("  getexp                  - Show expiration value")
	fmt.Println("  help | ?                - Show this help")

	fmt.Println("")
	fmt.Println("Datastructures supported:")
	fmt.Println("  hashmap - key/value store. Commands: create, set, get, del, keys/ls, clear, size, expire")
	fmt.Println("  array   - ordered list. Commands: insert, setindex, get (by index), push (append), pop, delete(index), ls, clear, size")
	fmt.Println("  stack   - LIFO. Commands: push, pop, peek, ls(size), clear, expire")
	fmt.Println("  queue   - FIFO. Commands: push (enqueue), dequeue/pop, peek, ls(size), clear, expire")
	fmt.Println("  minheap - priority queue. Commands: push <score> <value>, pop, peek, ls(size), clear, expire")
}
