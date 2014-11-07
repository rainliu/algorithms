package graph

import (
	"fmt"
	"os"
	"testing"
)

func TestUnigraph(t *testing.T) {
	fd, err := os.Open("tinyG.txt")
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", "tinyG.txt", err))
	}
	defer fd.Close()

	g := NewUnigraphFromReader(fd)

	fmt.Printf(g.String())
}

func TestDigraph(t *testing.T) {
	fd, err := os.Open("tinyDG.txt")
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", "tinyDG.txt", err))
	}
	defer fd.Close()

	g := NewDigraphFromReader(fd)

	fmt.Printf(g.String())
}
