package graph

import (
	"algorithms/container"
)

//Depth First Order
type DFO struct {
	marked      []bool
	pre         *container.Queue
	post        *container.Queue
	reversePost *container.Stack
}

func NewDFO(G *graph) *DFO {
	this := &DFO{}

	this.pre = &container.Queue{}
	this.post = &container.Queue{}
	this.reversePost = &container.Stack{}
	this.marked = make([]bool, G.V())

	for v := 0; v < G.V(); v++ {
		if !this.marked[v] {
			this.explore(G, v)
		}
	}

	return this
}

func (this *DFO) explore(G *graph, v int) {
	this.marked[v] = true

	this.previsit(v)
	iter := G.Adj(v).Iterator()
	for iter.HasNext() {
		w := iter.Next().Value.(Edge).To()
		if !this.marked[w] {
			this.explore(G, w)
		}
	}
	this.postvisit(v)
}

func (this *DFO) previsit(v int) {
	this.pre.Push(v)
}

func (this *DFO) postvisit(v int) {
	this.post.Push(v)
	this.reversePost.Push(v)
}

func (this *DFO) Pre() container.Iterable {
	return this.pre
}

func (this *DFO) Post() container.Iterable {
	return this.post
}

func (this *DFO) ReversePost() container.Iterable {
	return this.reversePost
}
