package graph

import (
	"algorithms/container"
)

type DFS struct {
	marked []bool
	edgeTo []int
	id     []int
	cc     int

	s int

	pre   []int
	post  []int
	clock int
}

func NewDFS(G *Graph, s int) *DFS {
	this := &DFS{}

	this.pre = make([]int, G.V())
	this.post = make([]int, G.V())
	this.clock = 0

	this.marked = make([]bool, G.V())
	this.edgeTo = make([]int, G.V())
	this.id = make([]int, G.V())
	this.s = s
	this.cc = 0

	for v := 0; v < G.V(); v++ {
		if !this.marked[v] {
			this.explore(G, v)
			this.cc++
		}
	}

	return this
}

func (this *DFS) explore(G *Graph, v int) {
	this.marked[v] = true

	this.previsit(v)
	iter := G.Adj(v).Iterator()
	for iter.HasNext() {
		w := iter.Next().Value.(int)
		if !this.marked[w] {
			this.edgeTo[w] = v
			this.explore(G, w)
		}
	}
	this.postvisit(v)
}

func (this *DFS) previsit(v int) {
	this.id[v] = this.cc
	this.pre[v] = this.clock
	this.clock++
}

func (this *DFS) postvisit(v int) {
	this.post[v] = this.clock
	this.clock++
}

func (this *DFS) Pre(v int) int {
	return this.pre[v]
}

func (this *DFS) Post(v int) int {
	return this.post[v]
}

//Paths interface
func (this *DFS) HasPathTo(v int) bool {
	return this.Connected(v, this.s)
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

//Search interface
func (this *DFS) Marked(v int) bool {
	return this.Connected(v, this.s)
}

//Search and CC interface
func (this *DFS) Count() int {
	return this.cc
}

//CC interface
func (this *DFS) Connected(v, w int) bool {
	return this.id[v] == this.id[w]
}

func (this *DFS) ID(v int) int {
	return this.id[v]
}
