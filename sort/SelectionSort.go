package sort

import (
	"algorithms"
)

type SelectionSort struct {
	sorter
}

func (this *SelectionSort) Sort(a []algorithms.Comparable) {
	N := len(a)
	for i := 0; i < N; i++ {
		min := i
		for j := i + 1; j < N; j++ {
			if this.Less(a[j], a[min]) {
				min = j
			}
		}
		this.Swap(a, i, min)
	}
}
