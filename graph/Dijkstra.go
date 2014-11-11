package graph

import (
	"algorithms/container"
	"math"
)

type Dijkstra struct {
	edgeTo []Edge
	distTo []float64
	pq     *MinHeap
}

func NewDijkstra(g *graph, s int) *Dijkstra {
	this := &Dijkstra{}

	this.edgeTo = make([]Edge, g.V())
	this.distTo = make([]float64, g.V())
	this.pq = container.NewMinHeap(g.V())

	for v := 0; v < g.V(); v++ {
		this.distTo[v] = math.MaxFloat64
	}
	this.distTo[s] = 0.0

	//this.pq.Push()

	return this
}
