package container

import ()

type Queue struct {
	container

	last *Item
}

func (this *Queue) Push(v interface{}) {
	oldlast := this.last
	this.last = &Item{}
	this.last.Value = v
	this.last.Next = nil

	if this.IsEmpty() {
		this.first = this.last
	} else {
		oldlast.Next = this.last
	}
	this.size++
}

func (this *Queue) Pop() *Item {
	item := this.first
	this.first = this.first.Next
	if this.IsEmpty() {
		this.last = nil
	}
	this.size--
	return item
}
