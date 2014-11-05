package graph

import ()

import (
	"algorithms/container"
	"fmt"
	"os"
	"testing"
)

func TestCC(t *testing.T) {
	fd, err := os.Open("tinyG.txt")
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", "tinyG.txt", err))
	}
	defer fd.Close()

	G := NewUnigraphFromReader(fd)

	var c CC
	c = NewDFS(G, 0)

	M := c.Count()
	fmt.Printf("%d components\n", M)

	components := make([]*container.Bag, M)
	for i := 0; i < M; i++ {
		components[i] = &container.Bag{}
	}
	for v := 0; v < G.V(); v++ {
		components[c.ID(v)].Push(v)
	}
	for i := 0; i < M; i++ {
		iter := components[i].Iterator()
		for iter.HasNext() {
			v := iter.Next().Value.(int)
			fmt.Printf("%d ", v)
		}
		fmt.Printf("\n")
	}
}
