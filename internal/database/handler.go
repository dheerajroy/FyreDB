package database

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/dheerajroy/fyredb/internal/datastructures"
)

type NodeType string

const (
	HashMapType NodeType = "hashmap"
	StackType   NodeType = "stack"
	QueueType   NodeType = "queue"
	ArrayType   NodeType = "array"
	MinHeapType NodeType = "minheap"
	ValueType   NodeType = "value"
)

type CurrentNode struct {
	Type  NodeType
	Value any
}

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
var nodeStack []CurrentNode
var nameStack []string

func Handle() {
	root := CurrentNode{Type: HashMapType, Value: &Store}
	nodeStack = []CurrentNode{root}
	nameStack = []string{"root"}

	for {
		fmt.Printf("fyredb (%s)> ", strings.Join(nameStack, "/"))
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		tokens := strings.Fields(line)
		if len(tokens) == 0 {
			continue
		}

		cmd := strings.ToLower(tokens[0])
		current := nodeStack[len(nodeStack)-1]

		switch cmd {
		case "exit", "quit":
			fmt.Println("Exiting fyredb.")
			return

		case "create":
			if len(tokens) < 3 {
				fmt.Println("Usage: create <name> <type>")
				continue
			}
			name := tokens[1]
			dsType := NodeType(tokens[2])

			if current.Type != HashMapType {
				fmt.Println("create is only supported inside hashmap nodes (databases)")
				continue
			}

			hm := current.Value.(*datastructures.HashMap)

			var unit *datastructures.Unit
			switch dsType {
			case HashMapType:
				unit = datastructures.NewUnit(datastructures.NewHashMap())
			case StackType:
				unit = datastructures.NewUnit(datastructures.NewStack())
			case QueueType:
				unit = datastructures.NewUnit(datastructures.NewQueue())
			case ArrayType:
				unit = datastructures.NewUnit(datastructures.NewArray())
			case MinHeapType:
				unit = datastructures.NewUnit(datastructures.NewMinHeap())
			default:
				fmt.Println("Unknown type:", dsType)
				continue
			}

			hm.Set(name, unit)
			fmt.Println("Created", dsType, name)

		case "use":
			if len(tokens) < 2 {
				fmt.Println("Usage: use <key_or_index>")
				continue
			}
			name := tokens[1]
			if !useNode(current, name) {
				continue
			}

		case "back":
			if len(nodeStack) <= 1 {
				fmt.Println("Already at root")
				continue
			}
			nodeStack = nodeStack[:len(nodeStack)-1]
			nameStack = nameStack[:len(nameStack)-1]
			fmt.Println("Moved back to parent node")

		case "root":
			nodeStack = []CurrentNode{root}
			nameStack = []string{"root"}
			fmt.Println("Now at root node")

		case "set":
			if len(tokens) < 3 {
				fmt.Println("Usage: set <key> <value>")
				continue
			}
			key := tokens[1]
			raw := tokens[2]

			if current.Type != HashMapType {
				fmt.Println("set is only supported inside hashmap nodes")
				continue
			}
			hm := current.Value.(*datastructures.HashMap)
			hm.Set(key, datastructures.NewUnit(parsePrimitive(raw)))
			fmt.Println("Set", key)

		case "get":
			if len(tokens) < 2 {
				fmt.Println("Usage: get <key_or_index>")
				continue
			}
			name := tokens[1]
			getNode(current, name)

		case "del":
			if len(tokens) < 2 {
				fmt.Println("Usage: del <key_or_index>")
				continue
			}
			name := tokens[1]
			delNode(current, name)

		case "ls":
			listNode(current)

		case "keys":
			listNode(current)

		case "clear":
			clearNode(current)

		case "size":
			sizeNode(current)

		case "insert":
			if len(tokens) < 3 {
				fmt.Println("Usage: insert <index> <value>")
				continue
			}
			if idx, err := strconv.Atoi(tokens[1]); err == nil {
				insertIndex(current, idx, tokens[2])
			} else {
				fmt.Println("Invalid index")
			}

		case "setindex":
			if len(tokens) < 3 {
				fmt.Println("Usage: setindex <index> <value>")
				continue
			}
			if idx, err := strconv.Atoi(tokens[1]); err == nil {
				setIndex(current, idx, tokens[2])
			} else {
				fmt.Println("Invalid index")
			}

		case "peek":
			peekNode(current)

		case "expire":
			if len(tokens) < 2 {
				fmt.Println("Usage: expire <milliseconds>")
				continue
			}
			if ms, err := strconv.ParseInt(tokens[1], 10, 64); err == nil {
				setExpiration(current, ms)
			} else {
				fmt.Println("Invalid expiration value")
			}

		case "getexp":
			getExpiration(current)

		case "push", "append":
			if len(tokens) < 2 {
				fmt.Println("Usage: push <value>")
				continue
			}

			if current.Type == MinHeapType && len(tokens) >= 3 {
				if score, err := strconv.ParseFloat(tokens[1], 64); err == nil {
					pushToMinHeap(current, score, tokens[2])
				} else {
					pushValue(current, tokens[1])
				}
				continue
			}
			raw := tokens[1]
			pushValue(current, raw)

		case "pop", "dequeue":
			popNode(current)

		case "help", "?":
			printHelp()

		default:
			fmt.Println("Unknown command:", cmd)
			fmt.Println("Available commands: create, use, back, root, set, get, del, ls, push, pop, exit")
		}
	}
}
