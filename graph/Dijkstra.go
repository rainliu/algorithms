package graph

import (
	"algorithms"
	"algorithms/container"
	"math"
)

type Dijkstra struct {
	edgeTo []Edge
	distTo []algorithms.Double
	pq     *container.PriorityQueue
	cycle  bool
}

func NewDijkstra(g *Graph, s int) *Dijkstra {
	this := &Dijkstra{}

	this.edgeTo = make([]Edge, g.V())
	this.distTo = make([]algorithms.Double, g.V())

	for v := 0; v < g.V(); v++ {
		this.distTo[v] = algorithms.Double(math.Inf(1))
	}
	this.distTo[s] = 0.0

	if NewCycle(g).HasCycle() {
		this.cycle = true
		this.pq = container.NewMinPriorityQueue(g.V())
		this.pq.Push(&container.PriorityQueueItem{s, algorithms.Double(0.0)})

		for !this.pq.IsEmpty() {
			pqi := this.pq.Pop().Value.(*container.PriorityQueueItem)
			this.relax(g, pqi.Index)
		}
	} else {
		this.cycle = false
		iter := NewDFO(g).ReversePost().Iterator()
		for iter.HasNext() {
			v := iter.Next().Value.(int)
			this.relax(g, v)
		}
	}

	return this
}

func (this *Dijkstra) relax(g *Graph, v int) {
	iter := g.Adj(v).Iterator()
	for iter.HasNext() {
		e := iter.Next().Value.(Edge)
		w := e.To()
		if this.distTo[w] > this.distTo[v]+e.Weight() {
			this.distTo[w] = this.distTo[v] + e.Weight()
			this.edgeTo[w] = e

			if this.cycle {
				if this.pq.Contains(w) {
					this.pq.ChangeKey(w, this.distTo[w])
				} else {
					this.pq.Push(&container.PriorityQueueItem{w, this.distTo[w]})
				}
			}
		}
	}
}

func (this *Dijkstra) DistTo(v int) algorithms.Double {
	return this.distTo[v]
}

func (this *Dijkstra) HasPathTo(v int) bool {
	return this.distTo[v] < algorithms.Double(math.Inf(1))
}

func (this *Dijkstra) PathTo(v int) container.Iterable {
	if !this.HasPathTo(v) {
		return nil
	}

	path := &container.Stack{}
	for e := this.edgeTo[v]; e != nil; e = this.edgeTo[e.From()] {
		path.Push(e)
	}
	return path
}
