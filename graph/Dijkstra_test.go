package graph

import (
	"fmt"
	"os"
	"testing"
)

func TestDijkstra(t *testing.T) {
	fd, err := os.Open("../data/tinyEWD.txt")
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", "tinyEWD.txt", err))
	}
	defer fd.Close()

	g := NewEWDigraphFromReader(fd)
	s := 0
	sp := NewDijkstra(g, s)

	for t := 0; t < g.V(); t++ {
		fmt.Printf("%d to %d", s, t)
		fmt.Printf(" (%4.2f): ", sp.DistTo(t))
		if sp.HasPathTo(t) {
			iter := sp.PathTo(t).Iterator()
			for iter.HasNext() {
				e := iter.Next().Value.(Edge)
				fmt.Printf("%s    ", e.String())
			}
		}
		fmt.Printf("\n")
	}
}

func TestAcyclicSP(t *testing.T) {
	fd, err := os.Open("../data/tinyEWDAG.txt")
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", "tinyEWDAG.txt", err))
	}
	defer fd.Close()

	g := NewEWDigraphFromReader(fd)
	s := 5
	sp := NewDijkstra(g, s)

	for t := 0; t < g.V(); t++ {
		fmt.Printf("%d to %d", s, t)
		fmt.Printf(" (%4.2f): ", sp.DistTo(t))
		if sp.HasPathTo(t) {
			iter := sp.PathTo(t).Iterator()
			for iter.HasNext() {
				e := iter.Next().Value.(Edge)
				fmt.Printf("%s    ", e.String())
			}
		}
		fmt.Printf("\n")
	}
}
