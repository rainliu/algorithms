package graph

import (
	"bufio"
	"io"
	"strings"
)

type SymbolGraph struct {
	g    *Graph
	keys []string
	st   map[string]int
}

func NewSymbolGraph(r io.Reader, delim string) *SymbolGraph {
	this := &SymbolGraph{}
	this.st = make(map[string]int)

	scanner1 := bufio.NewScanner(r)
	scanner1.Split(bufio.ScanLines)
	for scanner1.Scan() {
		strs := strings.Split(scanner1.Text(), delim)
		for i := 0; i < len(strs); i++ {
			if _, present := this.st[strs[i]]; !present {
				this.st[strs[i]] = len(this.st)
			}
		}
	}

	this.keys = make([]string, len(this.st))
	for name := range this.st {
		this.keys[this.st[name]] = name
	}

	this.g = NewGraph(len(this.st))

	scanner2 := bufio.NewScanner(r)
	scanner2.Split(bufio.ScanLines)
	for scanner2.Scan() {
		strs := strings.Split(scanner2.Text(), delim)
		v := this.st[strs[0]]
		for i := 1; i < len(strs); i++ {
			this.g.AddEdge(v, this.st[strs[i]])
		}
	}

	return this
}

func (this *SymbolGraph) Contains(s string) bool {
	_, present := this.st[s]

	return present
}

func (this *SymbolGraph) Index(s string) int {
	return this.st[s]
}

func (this *SymbolGraph) Name(v int) string {
	return this.keys[v]
}

func (this *SymbolGraph) G() *Graph {
	return this.g
}
