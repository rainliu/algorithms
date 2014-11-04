package graph

import (
	"algorithms/container"
)

type DFS struct {
	//Search interface
	marked []bool
	count  int

	//Paths interface
	edgeTo []int
	s      int

	//CC interface
	id []int
	cc int
}

func NewDFS(G *Graph, s int) *DFS {
	this := &DFS{}

	this.marked = make([]bool, G.V())
	this.edgeTo = make([]int, G.V())
	this.id = make([]int, G.V())

	this.s = s

	this.explore(G, s)

	return this
}

func NewCC(G *Graph) *DFS {
	this := &DFS{}

	this.marked = make([]bool, G.V())
	this.edgeTo = make([]int, G.V())
	this.id = make([]int, G.V())

	for s := 0; s < G.V(); s++ {
		if !this.marked[s] {
			this.explore(G, s)
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
	this.count++
	this.id[v] = this.cc
}

func (this *DFS) postvisit(v int) {

}

//Search interface
func (this *DFS) Marked(v int) bool {
	return this.marked[v]
}

func (this *DFS) Count() int {
	return this.count
}

//Paths interface
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

//CC interface
func (this *DFS) CC() int {
	return this.cc
}

func (this *DFS) Connected(v, w int) bool {
	return this.id[v] == this.id[w]
}

func (this *DFS) ID(v int) int {
	return this.id[v]
}
