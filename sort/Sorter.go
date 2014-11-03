package sort

import (
	"algorithms"
)

type Sorter interface {
	Less(v, w algorithms.Comparable) bool
	Swap(a []algorithms.Comparable, i, j int)
	IsSorted(a []algorithms.Comparable) bool

	Sort(a []algorithms.Comparable)
}

type sorter struct {
}

func (this *sorter) Less(v, w algorithms.Comparable) bool {
	return v.CompareTo(w) < 0
}

func (this *sorter) Swap(a []algorithms.Comparable, i, j int) {
	t := a[i]
	a[i] = a[j]
	a[j] = t
}

func (this *sorter) IsSorted(a []algorithms.Comparable) bool {
	for i := 1; i < len(a); i++ {
		if this.Less(a[i], a[i-1]) {
			return false
		}
	}

	return true
}
