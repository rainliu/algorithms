package sort

import (
	"algorithms"
    "testing"
    "algorithms/container"
    "fmt"
)

var b = []float64{664.0, 4121.85, 2678, 4409, 837.42, 3229.27, 4732.35, 4381.21, 66.10, 4747.08, 2156.86, 1025.70, 2520.97, 708.95, 3532.36, 4050.20}

func TestMaxHeap(t *testing.T) {
	M := 5
	a := algorithms.NewDoubleSlice(b)
	maxPQ := NewMaxHeap(M+1);
	for i:=0; i<len(a); i++ {
		maxPQ.Push(a[i])
		if maxPQ.Size() > M {
			maxPQ.Pop()
		}
	}
	
	var stack container.Stack
	for !maxPQ.IsEmpty() {
		stack.Push(maxPQ.Pop().Value.(algorithms.Double))
	}
	
	iter := stack.Iterator()
	for iter.HasNext() {
		fmt.Printf("%f\n", iter.Next().Value.(algorithms.Double))
	}	
}


func TestMinHeap(t *testing.T) {
	M := 5
	a := algorithms.NewDoubleSlice(b)
	minPQ := NewMinHeap(M+1);
	for i:=0; i<len(a); i++ {
		minPQ.Push(a[i])
		if minPQ.Size() > M {
			minPQ.Pop()
		}
	}
	
	var stack container.Stack
	for !minPQ.IsEmpty() {
		stack.Push(minPQ.Pop().Value.(algorithms.Double))
	}
	
	iter := stack.Iterator()
	for iter.HasNext() {
		fmt.Printf("%f\n", iter.Next().Value.(algorithms.Double))
	}
}
