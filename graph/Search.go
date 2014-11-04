package graph

import ()

type Search interface {
	Marked(v int) bool
	Count() int
}
