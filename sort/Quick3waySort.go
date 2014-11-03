package sort

import (
	"algorithms"
	"math/rand"
	"time"
)

type Quick3waySort struct {
	sorter
}

func (this *Quick3waySort) Sort(a []algorithms.Comparable) {
	this.sort(a, 0, len(a)-1)
}

func (this *Quick3waySort) sort(a []algorithms.Comparable, low, high int){
	if high<=low {
		return
	}

	//random swap low and x
	r:= rand.New(rand.NewSource(time.Now().UnixNano()))
	x:= low + (r.Int()%(high+1-low));
	this.Swap(a, low, x)
	
	/////////
	lt := low
	i  := low+1
	gt := high
	v  := a[low]
	
	for i<=gt {	
		cmp:= a[i].CompareTo(v);
		if cmp < 0 {
			this.Swap(a, lt, i)
			lt++
			i++
		}else if cmp > 0 {
			this.Swap(a, i, gt)
			gt--
		}else{
			i++
		}
	}

	this.sort(a, low, lt-1)
	this.sort(a, gt+1, high)
}
