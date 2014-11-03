package sort

import (
	"algorithms"
	"math/rand"
	"time"
)

type QuickSort struct {
	sorter
}

func (this *QuickSort) Sort(a []algorithms.Comparable) {
	this.sort(a, 0, len(a)-1)
}

func (this *QuickSort) sort(a []algorithms.Comparable, low, high int){
	if high<=low {
		return
	}
	
	j:= this.partition(a, low, high);
	this.sort(a, low, j-1)
	this.sort(a, j+1, high)
}

func (this *QuickSort) partition(a []algorithms.Comparable, low, high int) int{
	//random swap low and x
	r:= rand.New(rand.NewSource(time.Now().UnixNano()))
	x:= low + (r.Int()%(high+1-low));
	this.Swap(a, low, x)

	////////
	i:=low;
	j:=high+1
	v:=a[low]
	
	for {
		for i=i+1; this.Less(a[i], v); i++ {
			if i==high {
				break;
			}
		}
		
		for j=j-1; this.Less(v, a[j]); j-- {
			if j==low {
				break;
			}
		}
		
		if i>=j {
			break;
		}
		
		this.Swap(a, i, j)
	}
	this.Swap(a, low, j)
	
	return j;
}