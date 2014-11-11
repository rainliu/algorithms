package graph

import "fmt"

type Edge interface {
	Weight() float64
	From() int
	To() int
	String() string
}

type edge struct {
	v      int
	w      int
	weight float64
}

func NewEdge(v, w int) Edge {
	return &edge{v: v, w: w, weight: 1.0}
}

func NewWeightedEdge(v, w int, weight float64) Edge {
	return &edge{v: v, w: w, weight: weight}
}

func (this *edge) Weight() float64 {
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
