package graph

import (
	"algorithms/container"
)

type Paths interface {
	HasPathTo(v int) bool
	PathTo(v int) container.Iterable
}
