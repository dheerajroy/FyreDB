package datastructures

type BaseDataStructureInterface interface {
	SetExpiration(exp int64)
	GetExpiration() int64
	Size() int
	Clear()
}
