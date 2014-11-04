package graph

import ()

import (
	"fmt"
	"os"
	"testing"
)

func TestPath(t *testing.T) {
	fd, err := os.Open("tinyCG.txt")
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", "tinyCG.txt", err))
	}
	defer fd.Close()

	G := NewGraphFromReader(fd)
	s := 0
	var p Paths
	p = NewDFS(G, s)
	for v := 0; v < G.V(); v++ {
		fmt.Printf("%d to %d : ", s, v)
		if p.HasPathTo(v) {
			iter := p.PathTo(v).Iterator()
			for iter.HasNext() {
				x := iter.Next().Value.(int)
				if x == s {
					fmt.Printf("%d", x)
				} else {
					fmt.Printf("-%d", x)
				}
			}
		}
		fmt.Printf("\n")
	}
}
