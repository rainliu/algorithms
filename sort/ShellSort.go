package sort

import (
	"algorithms"
)

type ShellSort struct{
	sorter
}

func (this *ShellSort) Sort(a []algorithms.Comparable){
	N:=len(a)
	h:=1
	
	for h<N/3 {
		h = 3*h+1
	}
	
	for h>=1 {
		for i:=h; i<N; i++{
			for j:=i; j>=h && this.Less(a[j], a[j-h]); j-=h {
				this.Swap(a, j, j-h);
			}
		}
		h = h/3
	}
}