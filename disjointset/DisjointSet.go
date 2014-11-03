package disjointset

import (
	"algorithms/container"
	"fmt"
)

type UnionFind interface {
	Union(p, q int)
	Find(p int) int
}

type DisjointSet interface {
	UnionFind
	container.Iterable
	
	Connected(p, q int) bool
	Count() int
}

type disjointset struct {
	UnionFind

	id    []int
	count int

	current int
}

func (this *disjointset) Init(N int) {
	this.count = N
	this.id = make([]int, N)
	for i := 0; i < N; i++ {
		this.id[i] = i
	}
}

func (this *disjointset) Count() int {
	return this.count
}

func (this *disjointset) Connected(p, q int) bool {
	return this.UnionFind.Find(p) == this.UnionFind.Find(q)
}

func (this *disjointset) Iterator() container.Iterator {
	this.current = 0
	return this
}

func (this *disjointset) HasNext() bool {
	return this.current < len(this.id)
}

func (this *disjointset) Next() *container.Item {
	item := &container.Item{}

	item.Value = this.id[this.current]
	item.Next = nil

	this.current++

	return item
}

func summary(iterable container.Iterable) {
	iter := iterable.Iterator()
	for iter.HasNext() {
		item := iter.Next()
		fmt.Printf("%d ", item.Value.(int))
	}
}
