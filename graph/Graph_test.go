package graph

import (
	"fmt"
	"os"
	"testing"
)

func TestGraph(t *testing.T) {
	fd, err := os.Open("tinyG.txt")
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", "tinyG.txt", err))
	}
	defer fd.Close()

	g := NewUnigraphFromReader(fd)

	fmt.Printf(g.String())
}
