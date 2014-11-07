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

	sg := NewSymbolGraph(fd, false, delim)
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

func TestDegreeOfSeparation(t *testing.T) {
	CallDegreeOfSeparation("routes.txt", "routes_cmd.txt", " ", "JFK")
}

func CallDegreeOfSeparation(filename, command, delim, source string) {
	fd, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filename, err))
	}
	defer fd.Close()

	sg := NewSymbolGraph(fd, false, delim)
	g := sg.G()
	//fmt.Printf("%s\n", g.String())

	if !sg.Contains(source) {
		fmt.Printf("%s not in database\n", source)
		return
	}

	s := sg.Index(source)
	bfs := NewBFS(g, s)

	cmd, err2 := os.Open(command)
	if err2 != nil {
		panic(fmt.Sprintf("open %s: %v", command, err2))
	}
	defer cmd.Close()

	scanner := bufio.NewScanner(cmd)
	for scanner.Scan() {
		sink := scanner.Text()
		fmt.Printf("%s\n", sink)
		if sg.Contains(sink) {
			t := sg.Index(sink)
			if bfs.HasPathTo(t) {
				iter := bfs.PathTo(t).Iterator()
				for iter.HasNext() {
					v := iter.Next().Value.(int)
					fmt.Printf("   %s\n", sg.Name(v))
				}
			} else {
				fmt.Printf("Not connected\n")
			}
		} else {
			fmt.Printf("Not in database\n")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
