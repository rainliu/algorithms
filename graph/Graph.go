package graph

import (
	"algorithms"
	"algorithms/container"
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type GraphType int

const (
	UNIGRAPH  GraphType = 0
	DIGRAPH   GraphType = 1
	EWDIGRAPH GraphType = 2
)

type Graph struct {
	gt  GraphType
	v   int
	e   int
	adj []*container.Bag
}

func NewUnigraph(v int) *Graph {
	return NewGraph(v, UNIGRAPH)
}

func NewDigraph(v int) *Graph {
	return NewGraph(v, DIGRAPH)
}

func NewEWDigraph(v int) *Graph {
	return NewGraph(v, EWDIGRAPH)
}

func NewGraph(v int, gt GraphType) *Graph {
	this := &Graph{}
	this.gt = gt
	this.v = v
	this.e = 0
	this.adj = make([]*container.Bag, v)
	for i := 0; i < v; i++ {
		this.adj[i] = &container.Bag{}
	}
	return this
}

func NewUnigraphFromReader(r io.Reader) *Graph {
	return NewGraphFromReader(r, UNIGRAPH)
}

func NewDigraphFromReader(r io.Reader) *Graph {
	return NewGraphFromReader(r, DIGRAPH)
}

func NewEWDigraphFromReader(r io.Reader) *Graph {
	return NewGraphFromReader(r, EWDIGRAPH)
}

func NewGraphFromReader(r io.Reader, gt GraphType) *Graph {
	var v, e, w int
	var err error
	var this *Graph

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	if v, err = strconv.Atoi(scanner.Text()); err != nil {
		fmt.Printf(err.Error())
		return nil
	}

	switch gt {
	case DIGRAPH:
		this = NewDigraph(v)
	case EWDIGRAPH:
		this = NewEWDigraph(v)
	default:
		this = NewUnigraph(v)
	}

	scanner.Scan()
	if e, err = strconv.Atoi(scanner.Text()); err != nil {
		fmt.Printf(err.Error())
		return nil
	}

	for i := 0; i < e; i++ {
		scanner.Scan()
		if v, err = strconv.Atoi(scanner.Text()); err != nil {
			fmt.Printf(err.Error())
			return nil
		}
		scanner.Scan()
		if w, err = strconv.Atoi(scanner.Text()); err != nil {
			fmt.Printf(err.Error())
			return nil
		}
		if gt == EWDIGRAPH {
			var weight float64
			scanner.Scan()
			if weight, err = strconv.ParseFloat(scanner.Text(), 64); err != nil {
				fmt.Printf(err.Error())
				return nil
			}
			this.AddWeightedEdge(v, w, algorithms.Double(weight))
		} else {
			this.AddEdge(v, w)
		}
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
	this.adj[v].Push(NewEdge(v, w))
	if this.gt == UNIGRAPH {
		this.adj[w].Push(NewEdge(w, v))
	}
	this.e++
}

func (this *Graph) AddWeightedEdge(v, w int, weight algorithms.Double) {
	this.adj[v].Push(NewWeightedEdge(v, w, weight))
	if this.gt == UNIGRAPH {
		this.adj[w].Push(NewWeightedEdge(w, v, weight))
	}
	this.e++
}

func (this *Graph) GraphType() GraphType {
	return this.gt
}

func (this *Graph) Reverse() *Graph {
	if this.gt != UNIGRAPH {
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
			w := iter.Next().Value.(Edge).To()
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
			w := iter.Next().Value.(Edge).To()
			s += strconv.Itoa(w) + " "
		}
		s += "\n"
	}
	return s
}
