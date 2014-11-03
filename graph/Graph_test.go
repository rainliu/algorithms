package graph


import (
    "testing"
    "fmt"
    "os"
)

func TestGraph(t *testing.T) {
	fd, err := os.Open("tinyG.txt")
    if err != nil {
        panic(fmt.Sprintf("open %s: %v", "tinyG.txt", err))
    }
    defer fd.Close()

    g := NewGraphFromReader(fd);

    fmt.Printf(g.String())
}