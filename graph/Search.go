package graph

import ()

type Search interface {
	HasPathTo(v int) bool
	Count() int
}
