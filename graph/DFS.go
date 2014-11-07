package graph

import (
	"algorithms/container"
)

//Depth First Search
type DFS struct {
	marked []bool
	edgeTo []int
	id     []int

	digraph bool
	count   int

	s int
}

func NewDFS(G *Graph, s int) *DFS {
	this := &DFS{}

	this.marked = make([]bool, G.V())
	this.edgeTo = make([]int, G.V())
	this.id = make([]int, G.V())
	this.s = s

	this.digraph = G.IsDigraph()
	this.count = 0

	if this.digraph {
		dfo := NewDFO(G.Reverse())

		iter := dfo.ReversePost().Iterator()
		for iter.HasNext() {
			v := iter.Next().Value.(int)
			if !this.marked[v] {
				this.explore(G, v, true)
				this.count++
			}
		}

		//reset marked slice
		this.marked = make([]bool, G.V())
		this.explore(G, s, false)
	} else {
		for v := 0; v < G.V(); v++ {
			if !this.marked[v] {
				this.explore(G, v, true)
				this.count++
			}
		}
	}
	return this
}

func (this *DFS) explore(G *Graph, v int, cc bool) {
	if cc {
		this.id[v] = this.count
	}
	this.marked[v] = true

	iter := G.Adj(v).Iterator()
	for iter.HasNext() {
		w := iter.Next().Value.(int)
		if !this.marked[w] {
			if !cc {
				this.edgeTo[w] = v
			}
			this.explore(G, w, cc)
		}
	}
}

//Paths interface
func (this *DFS) HasPathTo(v int) bool {
	if this.digraph {
		return this.marked[v]
	} else {
		return this.Connected(v, this.s)
	}
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

//SC/SCC interface
func (this *DFS) Count() int {
	return this.count
}

//CC/SCC interface
func (this *DFS) Connected(v, w int) bool {
	return this.id[v] == this.id[w]
}

func (this *DFS) ID(v int) int {
	return this.id[v]
}
