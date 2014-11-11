package graph

import (
	"algorithms/container"
)

//Breadth First Search
type BFS struct {
	marked []bool
	edgeTo []int
	s      int
}

func NewBFS(G *graph, s int) *BFS {
	this := &BFS{}

	this.marked = make([]bool, G.V())
	this.edgeTo = make([]int, G.V())
	this.s = s

	this.explore(G, s)

	return this
}

func (this *BFS) explore(G *graph, s int) {
	this.marked[s] = true

	queue := &container.Queue{}
	queue.Push(s)
	for !queue.IsEmpty() {
		v := queue.Pop().Value.(int)
		iter := G.Adj(v).Iterator()
		for iter.HasNext() {
			w := iter.Next().Value.(Edge).To()
			if !this.marked[w] {
				this.edgeTo[w] = v
				this.marked[w] = true
				queue.Push(w)
			}
		}
	}
}

func (this *BFS) HasPathTo(v int) bool {
	return this.marked[v]
}

func (this *BFS) PathTo(v int) container.Iterable {
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
