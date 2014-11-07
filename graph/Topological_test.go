package graph

import ()

import (
	"fmt"
	"os"
	"testing"
)

func TestTopological(t *testing.T) {
	fd, err := os.Open("jobs.txt")
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", "jobs.txt", err))
	}
	defer fd.Close()

	sg := NewSymbolGraph(fd, true, "/")
	G := sg.G()

	cycleFinder := NewCycle(G)
	if !cycleFinder.HasCycle() {
		dfo := NewDFO(G)
		iter := dfo.ReversePost().Iterator()
		for iter.HasNext() {
			v := iter.Next().Value.(int)
			fmt.Printf("%s\n", sg.Name(v))
		}
	} else {
		fmt.Printf("Cycle Detected:\n")
		iter := cycleFinder.Cycle().Iterator()
		for iter.HasNext() {
			v := iter.Next().Value.(int)
			fmt.Printf("  %s\n", sg.Name(v))
		}
	}
}
