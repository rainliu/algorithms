package graph

import (
	"algorithms"
	"fmt"
)

type Edge interface {
	Weight() algorithms.Double
	From() int
	To() int
	String() string
}

type edge struct {
	v      int
	w      int
	weight algorithms.Double
}

func NewEdge(v, w int) Edge {
	return &edge{v: v, w: w, weight: 1.0}
}

func NewWeightedEdge(v, w int, weight algorithms.Double) Edge {
	return &edge{v: v, w: w, weight: weight}
}

func (this *edge) Weight() algorithms.Double {
	return this.weight
}

func (this *edge) From() int {
	return this.v
}

func (this *edge) To() int {
	return this.w
}

func (this *edge) String() string {
	return fmt.Sprintf("%d->%d %.2f", this.v, this.w, this.weight)
}
