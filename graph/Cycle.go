package graph

import (
	"algorithms/container"
)

type Cycle struct {
	hasCycle bool
	marked   []bool
	edgeTo   []int

	cycle   *container.Stack
	onStack []bool
}

func NewCycle(G *graph) *Cycle {
	this := &Cycle{}

	this.hasCycle = false
	this.cycle = nil
	this.marked = make([]bool, G.V())
	this.edgeTo = make([]int, G.V())

	if G.IsDigraph() {
		this.onStack = make([]bool, G.V())
	}

	for s := 0; s < G.V(); s++ {
		if !this.marked[s] {
			this.explore(G, s, s)
		}
	}

	return this
}

func (this *Cycle) explore(G *graph, v, u int) {
	if G.IsDigraph() {
		this.onStack[v] = true
	}
	this.marked[v] = true
	iter := G.Adj(v).Iterator()
	for iter.HasNext() {
		w := iter.Next().Value.(Edge).To()
		if this.HasCycle() {
			return
		} else if !this.marked[w] {
			this.edgeTo[w] = v
			this.explore(G, w, v)
		} else {
			if G.IsDigraph() {
				if this.onStack[w] {
					this.hasCycle = true
				}
			} else {
				if w != u {
					this.hasCycle = true
				}
			}
			if this.hasCycle {
				this.cycle = &container.Stack{}
				for x := v; x != w; x = this.edgeTo[x] {
					this.cycle.Push(x)
				}
				this.cycle.Push(w)
				this.cycle.Push(v)
			}
		}
	}
	if G.IsDigraph() {
		this.onStack[v] = false
	}
	return
}

func (this *Cycle) HasCycle() bool {
	return this.hasCycle
}

func (this *Cycle) Cycle() container.Iterable {
	return this.cycle
}
