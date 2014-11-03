package graph

import (
	"algorithms/container"
	"io"
)

type Graph struct{
	v int
	e int
	adj []*container.Bag
}

func NewGraph(v int)*Graph{
	this := &Graph{}
	this.v = v;
	this.e = 0;
	this.adj = make([]*container.Bag, v);
	for i:=0; i<v; i++{
		this.adj[i] = &container.Bag{};
	}
	return this
}

func NewGraphFromReader(r io.Reader)*Graph{
	this := &Graph{}

	return this
}

func (this *Graph) V() int{
	return this.v;
}

func (this *Graph) E() int{
	return this.e;
}

func (this *Graph) AddEdge(v, w int){
	this.adj[v].Push(w)
	this.adj[w].Push(v)
	this.e++
}

func (this *Graph) Adj(v int) container.Iterable{
	return this.adj[v]
}