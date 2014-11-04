package graph

import (
	"algorithms/container"
)

type DFS struct {
	marked []bool
	count  int

	s      int
	edgeTo []int
}

func NewDFS(G *Graph, s int) *DFS {
	this := &DFS{}

	this.marked = make([]bool, G.V())
	this.count = 0
	this.s = s
	this.edgeTo = make([]int, G.V())

	this.explore(G, s)

	return this
}

func (this *DFS) explore(G *Graph, v int) {
	this.marked[v] = true

	this.previsit()
	iter := G.Adj(v).Iterator()
	for iter.HasNext() {
		w := iter.Next().Value.(int)
		if !this.marked[w] {
			this.edgeTo[w] = v
			this.explore(G, w)
		}
	}
	this.postvisit()
}

func (this *DFS) previsit() {
	this.count++
}

func (this *DFS) postvisit() {

}

func (this *DFS) Marked(v int) bool {
	return this.marked[v]
}

func (this *DFS) Count() int {
	return this.count
}

func (this *DFS) HasPathTo(v int) bool {
	return this.marked[v]
}

func (this *DFS) PathTo(v int) container.Iterable {
	if !this.HasPathTo(v) {
		return nil
	}

	path := &container.Stack{}
	for x := v; x != this.s; x = this.edgeTo[x] {
		path.Push(x)
	}
	path.Push(this.s)
	return path
}
