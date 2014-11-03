package sort

import (
	"algorithms"
)

type HeapSort struct {
	sorter
}

func (this *HeapSort) Sort(a []algorithms.Comparable) {
	maxPQ:=NewMaxHeap(len(a));
	maxPQ.N = len(a)
	
	for i:=0; i<len(a); i++{
		maxPQ.pq[i+1] = a[i]
	}
	
	for k := maxPQ.N/2; k >= 1; k-- {
		maxPQ.sink(maxPQ.pq, k, maxPQ.N)
	}

	for maxPQ.N > 1 {
		maxPQ.swap(maxPQ.pq, 1, maxPQ.N)
		maxPQ.N--
		maxPQ.sink(maxPQ.pq, 1, maxPQ.N)
	}
	
	for i:=0; i<len(a); i++{
		a[i] = maxPQ.pq[i+1]
	}
}
