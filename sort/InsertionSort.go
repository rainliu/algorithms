package sort

import (
	"algorithms"
)

type InsertionSort struct{
	sorter
}

func (this *InsertionSort) Sort(a []algorithms.Comparable){
	N:=len(a)
	
	for i:=1; i<N; i++ {
		for j:=i; j>0 && this.Less(a[j], a[j-1]); j-- {
			this.Swap(a, j, j-1);
		}
	}
}