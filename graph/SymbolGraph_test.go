package graph

import ()

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestSymbolGraph(t *testing.T) {
	CallSymbolGraph("routes.txt", "routes_cmd.txt", " ")
	CallSymbolGraph("movies.txt", "movies_cmd.txt", "/")
}

func CallSymbolGraph(filename, command, delim string) {
	fd, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filename, err))
	}
	defer fd.Close()

	sg := NewSymbolGraph(fd, delim)
	g := sg.G()
	//fmt.Printf("%s\n", g.String())

	cmd, err2 := os.Open(command)
	if err2 != nil {
		panic(fmt.Sprintf("open %s: %v", command, err2))
	}
	defer cmd.Close()

	scanner := bufio.NewScanner(cmd)
	for scanner.Scan() {
		source := scanner.Text()
		fmt.Printf("%s\n", source)
		iter := g.Adj(sg.Index(source)).Iterator()
		for iter.HasNext() {
			w := iter.Next().Value.(int)
			fmt.Printf("   %s\n", sg.Name(w))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
