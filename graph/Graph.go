package graph

import (
	"algorithms/container"
	"bufio"
	"io"
	"strconv"
)

type Graph struct {
	v   int
	e   int
	adj []*container.Bag
}

func NewGraph(v int) *Graph {
	this := &Graph{}
	this.v = v
	this.e = 0
	this.adj = make([]*container.Bag, v)
	for i := 0; i < v; i++ {
		this.adj[i] = &container.Bag{}
	}
	return this
}

func NewGraphFromReader(r io.Reader) *Graph {
	var v, e, w int
	var err error

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	if v, err = strconv.Atoi(scanner.Text()); err != nil {
		return nil
	}

	this := NewGraph(v)

	scanner.Scan()
	if e, err = strconv.Atoi(scanner.Text()); err != nil {
		return nil
	}

	for i := 0; i < e; i++ {
		scanner.Scan()
		if v, err = strconv.Atoi(scanner.Text()); err != nil {
			return nil
		}
		scanner.Scan()
		if w, err = strconv.Atoi(scanner.Text()); err != nil {
			return nil
		}
		this.AddEdge(v, w)
	}
	return this
}

func (this *Graph) V() int {
	return this.v
}

func (this *Graph) E() int {
	return this.e
}

func (this *Graph) AddEdge(v, w int) {
	this.adj[v].Push(w)
	this.adj[w].Push(v)
	this.e++
}

func (this *Graph) Adj(v int) container.Iterable {
	return this.adj[v]
}

func (this *Graph) Degree(v int) int {
	degree := 0
	iter := this.adj[v].Iterator()
	for iter.HasNext() {
		iter.Next()
		degree++
	}
	return degree
}

func (this *Graph) MaxDegree() int {
	max := 0
	for v := 0; v < this.v; v++ {
		degree := this.Degree(v)
		if degree > max {
			max = degree
		}
	}
	return max
}

func (this *Graph) AvgDegree() int {
	return 2 * this.e / this.v
}

func (this *Graph) NumberOfSelfLoops() int {
	count := 0
	for v := 0; v < this.v; v++ {
		iter := this.adj[v].Iterator()
		for iter.HasNext() {
			item := iter.Next()
			w := item.Value.(int)
			if v == w {
				count++
			}
		}
	}
	return count / 2
}

func (this *Graph) String() string {
	var s string
	s = strconv.Itoa(this.v) + " vertices, " + strconv.Itoa(this.e) + " edges\n"
	for v := 0; v < this.v; v++ {
		s += strconv.Itoa(v) + ": "
		iter := this.adj[v].Iterator()
		for iter.HasNext() {
			item := iter.Next()
			w := item.Value.(int)
			s += strconv.Itoa(w) + " "
		}
		s += "\n"
	}
	return s
}
