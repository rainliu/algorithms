package container

import ()

type Bag struct {
	container
}

func (this *Bag) Push(v interface{}) {
	oldfirst := this.first
	this.first = &Item{}
	this.first.Value = v
	this.first.Next = oldfirst
	this.size++
}

func (this *Bag) Pop() *Item {
	return nil
}
