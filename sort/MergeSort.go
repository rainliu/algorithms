package sort

import (
	"algorithms"
)

type MergeSort struct {
	sorter

	aux []algorithms.Comparable
}

func (this *MergeSort) Sort(a []algorithms.Comparable) {
	this.aux = make([]algorithms.Comparable, len(a))

	this.sort(a, 0, len(a)-1)
}

func (this *MergeSort) sort(a []algorithms.Comparable, low, high int) {
	if high<=low {
		return
	}
	
	mid := low + (high-low)/2
	
	this.sort(a, low, mid)
	this.sort(a, mid+1, high)
	this.merge(a, low, mid, high)
}

func (this *MergeSort) merge(a []algorithms.Comparable, low, mid, high int) {
	i := low
	j := mid + 1

	for k := low; k <= high; k++ {
		this.aux[k] = a[k]
	}

	for k := low; k <= high; k++ {
		if i > mid {
			a[k] = this.aux[j]
			j++
		} else if j > high {
			a[k] = this.aux[i]
			i++
		} else if this.Less(this.aux[j], this.aux[i]) {
			a[k] = this.aux[j]
			j++
		} else {
			a[k] = this.aux[i]
			i++
		}
	}
}
