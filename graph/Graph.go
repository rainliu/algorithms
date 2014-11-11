package graph

import (
	"algorithms/container"
	"bufio"
	"io"
	"strconv"
)

type graph struct {
	digraph bool
	v       int
	e       int
	adj     []*container.Bag
}

func NewUnigraph(v int) *graph {
	return NewGraph(v, false)
}

func NewDigraph(v int) *graph {
	return NewGraph(v, true)
}

func NewUnigraphFromReader(r io.Reader) *graph {
	return NewGraphFromReader(r, false)
}

func NewDigraphFromReader(r io.Reader) *graph {
	return NewGraphFromReader(r, true)
}

func NewGraph(v int, digraph bool) *graph {
	this := &graph{}
	this.digraph = digraph
	this.v = v
	this.e = 0
	this.adj = make([]*container.Bag, v)
	for i := 0; i < v; i++ {
		this.adj[i] = &container.Bag{}
	}
	return this
}

func NewGraphFromReader(r io.Reader, digraph bool) *graph {
	var v, e, w int
	var err error
	var this *graph

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	if v, err = strconv.Atoi(scanner.Text()); err != nil {
		return nil
	}

	if digraph {
		this = NewDigraph(v)
	} else {
		this = NewUnigraph(v)
	}

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

func (this *graph) V() int {
	return this.v
}

func (this *graph) E() int {
	return this.e
}

func (this *graph) AddEdge(v, w int) {
	this.adj[v].Push(NewEdge(v, w))
	if !this.digraph {
		this.adj[w].Push(NewEdge(w, v))
	}
	this.e++
}

func (this *graph) AddWeightedEdge(v, w int, weight float64) {
	this.adj[v].Push(NewWeightedEdge(v, w, weight))
	if !this.digraph {
		this.adj[w].Push(NewWeightedEdge(w, v, weight))
	}
	this.e++
}

func (this *graph) IsDigraph() bool {
	return this.digraph
}

func (this *graph) Reverse() *graph {
	if this.digraph {
		R := NewDigraph(this.v)
		for v := 0; v < this.v; v++ {
			iter := this.Adj(v).Iterator()
			for iter.HasNext() {
				w := iter.Next().Value.(Edge).To()
				R.AddEdge(w, v)
			}
		}
		return R
	} else {
		return nil
	}
}

func (this *graph) Adj(v int) container.Iterable {
	return this.adj[v]
}

func (this *graph) Degree(v int) int {
	degree := 0
	iter := this.adj[v].Iterator()
	for iter.HasNext() {
		iter.Next()
		degree++
	}
	return degree
}

func (this *graph) MaxDegree() int {
	max := 0
	for v := 0; v < this.v; v++ {
		degree := this.Degree(v)
		if degree > max {
			max = degree
		}
	}
	return max
}

func (this *graph) AvgDegree() int {
	return 2 * this.e / this.v
}

func (this *graph) NumberOfSelfLoops() int {
	count := 0
	for v := 0; v < this.v; v++ {
		iter := this.adj[v].Iterator()
		for iter.HasNext() {
			w := iter.Next().Value.(Edge).To()
			if v == w {
				count++
			}
		}
	}
	return count / 2
}

func (this *graph) String() string {
	var s string
	s = strconv.Itoa(this.v) + " vertices, " + strconv.Itoa(this.e) + " edges\n"
	for v := 0; v < this.v; v++ {
		s += strconv.Itoa(v) + ": "
		iter := this.adj[v].Iterator()
		for iter.HasNext() {
			w := iter.Next().Value.(Edge).To()
			s += strconv.Itoa(w) + " "
		}
		s += "\n"
	}
	return s
}
