package sort

import (
	"algorithms"
	"algorithms/container"
)

type priorityqueue interface{
	comp(a []algorithms.Comparable, i,j int) bool
}

type heap struct{
	priorityqueue
	
	pq []algorithms.Comparable
	N	int
}

func (this *heap) Init(maxN int) {
	this.pq = make([]algorithms.Comparable, maxN+1)
	this.N = 0;
}

func (this *heap) IsEmpty() bool{
	return this.N==0
}

func (this *heap) Size() int{
	return this.N
}

func (this *heap) Push(v interface{}){
	this.N++
	if c, ok := v.(algorithms.Comparable); ok{
		this.pq[this.N] = c;
		this.swim(this.pq, this.N)
	}else{
		panic("v is not Comparable interface!\n")
	}
}

func (this *heap) Pop() *container.Item{
	item := &container.Item{}
	item.Next = nil
	item.Value = this.pq[1];
	
	this.swap(this.pq, 1, this.N)
	this.N--
	this.pq[this.N+1] = nil
	this.sink(this.pq, 1, this.N)
	return item;
}

func (this *heap) swap(a []algorithms.Comparable, i, j int){
	t:=a[i]
	a[i] = a[j]
	a[j] = t
}

func (this *heap) sink(a []algorithms.Comparable, k, N int){
	for 2*k <= N {
		j:=2*k
		if j<N && this.priorityqueue.comp(a, j, j+1) {
			j++;
		}
		if !this.priorityqueue.comp(a, k, j) {
			break;
		}
		this.swap(a, k,j)
		k=j
	}
}

func (this *heap) swim(a []algorithms.Comparable, k int){
	for k>1 && this.priorityqueue.comp(a, k/2, k) {
		this.swap(a, k/2, k)
		k=k/2
	}
}

type MaxHeap struct{
	heap
}

func NewMaxHeap(maxN int) *MaxHeap{
	this := &MaxHeap{}
	this.priorityqueue = this
	this.Init(maxN);
	return this;
}

func (this *MaxHeap) comp(a []algorithms.Comparable, i, j int) bool{
	return a[i].CompareTo(a[j])<0;
}


type MinHeap struct{
	heap
}

func NewMinHeap(maxN int) *MinHeap{
	this := &MinHeap{}
	this.priorityqueue = this
	this.Init(maxN);
	return this;
}

func (this *MinHeap) comp(a []algorithms.Comparable, i, j int) bool{
	return a[i].CompareTo(a[j])>0;
}
