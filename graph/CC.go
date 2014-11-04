package graph

import ()

type CC interface {
	CC() int
	Connected(v, w int) bool
	ID(v int) int
}
