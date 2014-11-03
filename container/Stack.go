package container

import ()

type Stack struct {
	container
}

func (this *Stack) Push(v interface{}) {
	oldfirst := this.first
	this.first = &Item{}
	this.first.Value = v
	this.first.Next = oldfirst
	this.size++
}

func (this *Stack) Pop() *Item {
	item := this.first
	this.first = this.first.Next
	this.size--
	return item
}
