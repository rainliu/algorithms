package search

import (
	"algorithms"
	"algorithms/container"
)

type SymbolTable interface{
	Set(key algorithms.Comparable, value interface{})
	Get(key algorithms.Comparable) interface{}
	Delete(key algorithms.Comparable)
	Contains(key algorithms.Comparable) bool
	IsEmpty() bool
	Size() int
	Min() algorithms.Comparable
	Max() algorithms.Comparable
	Floor(algorithms.Comparable) algorithms.Comparable
	Ceil(algorithms.Comparable) algorithms.Comparable
	Rank(key algorithms.Comparable) int
	Select(k int) algorithms.Comparable
	DeleteMin() 
	DeleteMax()
	RangedSize(low, high algorithms.Comparable) int
	RangedKeys(low, high algorithms.Comparable) container.Iterator
	Keys() container.Iterator
}
