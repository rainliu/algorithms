package graph

import ()

import (
	"fmt"
	"os"
	"testing"
)

func TestUnigraphSearch(t *testing.T) {
	fd, err := os.Open("tinyG.txt")
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", "tinyG.txt", err))
	}
	defer fd.Close()

	G := NewUnigraphFromReader(fd)

	var s Search

	s = NewDFS(G, 0)
	for v := 0; v < G.V(); v++ {
		if s.HasPathTo(v) {
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
		if s.HasPathTo(v) {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Printf("\n")
	if s.Count() != 1 {
		fmt.Printf("NOT ")
	}
	fmt.Printf("connected\n")
}

func TestDigraphSearch(t *testing.T) {
	fd, err := os.Open("tinyDG.txt")
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", "tinyDG.txt", err))
	}
	defer fd.Close()

	G := NewDigraphFromReader(fd)

	var s Search

	s = NewDFS(G, 1)
	for v := 0; v < G.V(); v++ {
		if s.HasPathTo(v) {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Printf("\n")

	s = NewDFS(G, 2)
	for v := 0; v < G.V(); v++ {
		if s.HasPathTo(v) {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Printf("\n")
}
