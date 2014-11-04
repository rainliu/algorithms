package graph

import ()

import (
	"fmt"
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	fd, err := os.Open("tinyG.txt")
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", "tinyG.txt", err))
	}
	defer fd.Close()

	G := NewGraphFromReader(fd)

	var s Search

	s = NewDFS(G, 0)
	for v := 0; v < G.V(); v++ {
		if s.Marked(v) {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Printf("\n")
	if s.Count() != 1 {
		fmt.Printf("NOT ")
	}
	fmt.Printf("connected\n")

	s = NewDFS(G, 9)
	for v := 0; v < G.V(); v++ {
		if s.Marked(v) {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Printf("\n")
	if s.Count() != 1 {
		fmt.Printf("NOT ")
	}
	fmt.Printf("connected\n")

}
