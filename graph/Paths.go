package graph

import (
	"algorithms/container"
)

type Paths interface {
	HasPathTo(v int) bool
	PathTo(v int) container.Iterable
}

type ShortestPaths interface {
	Paths
	DistTo(v int) float64
}
