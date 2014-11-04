package graph

import ()

type CC interface {
	Count() int
	Connected(v, w int) bool
	ID(v int) int
}
