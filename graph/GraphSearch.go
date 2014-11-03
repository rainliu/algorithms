package graph

import ()

type GraphSearch interface {
	Marked(v int) bool
	Count() int
}
