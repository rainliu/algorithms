package container

import ()

type Iterator interface {
	HasNext() bool
	Next() *Item
}

type Iterable interface{
	Iterator() Iterator
}

type Container interface {
	Iterable
	
	IsEmpty() bool
	Size() int
	Push(v interface{})
	Pop() *Item
}

type Item struct {
	Value interface{}

	Next *Item
}

type container struct {
	first *Item
	size  int

	current *Item
}

func (this *container) IsEmpty() bool {
	return this.first == nil
}

func (this *container) Size() int {
	return this.size
}

func (this *container) Iterator() Iterator {
	this.current = this.first
	return this
}

func (this *container) HasNext() bool {
	return this.current != nil
}

func (this *container) Next() *Item {
	item := this.current
	this.current = this.current.Next
	return item
}
