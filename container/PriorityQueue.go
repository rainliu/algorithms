package container

import (
	"algorithms"
)

type PriorityQueueItem struct {
	Index int
	Key   algorithms.Comparable
}

type priority_queue struct {
	max_pq  bool
	current int
	size    int
	pq      []int
	qp      []int
	keys    []algorithms.Comparable
}

func NewMaxPriorityQueue(size int) *priority_queue {
	return NewPriorityQueue(size, true)
}

func NewMinPriorityQueue(size int) *priority_queue {
	return NewPriorityQueue(size, false)
}

func NewPriorityQueue(size int, max_pq bool) *priority_queue {
	this := &priority_queue{}

	this.pq = make([]int, size+1)
	this.qp = make([]int, size+1)
	this.keys = make([]algorithms.Comparable, size+1)
	this.size = 0
	this.current = 0
	this.max_pq = max_pq

	for i := 0; i <= size; i++ {
		this.qp[i] = -1
	}

	return this
}

func (this *priority_queue) IsEmpty() bool {
	return this.size == 0
}

func (this *priority_queue) Size() int {
	return this.size
}

func (this *priority_queue) Push(v interface{}) {
	if pqi, ok := v.(*PriorityQueueItem); ok {
		if this.Contains(pqi.Index) {
			panic("index is already in the priority queue\n")
		} else {
			this.size++

			this.qp[pqi.Index] = this.size
			this.pq[this.size] = pqi.Index
			this.keys[pqi.Index] = pqi.Key

			this.swim(this.size)
		}
	} else {
		panic("input v is not *PriorityQueueItem\n")
	}
}

func (this *priority_queue) Pop() *Item {
	if this.size == 0 {
		panic("Priority queue underflow\n")
	}

	idx := this.pq[1]

	item := &Item{}
	item.Next = nil
	item.Value = &PriorityQueueItem{this.pq[1], this.keys[this.pq[1]]}

	this.swap(1, this.size)
	this.size--
	this.sink(1)

	this.qp[idx] = -1
	this.keys[this.pq[this.size+1]] = nil
	this.pq[this.size+1] = -1

	return item
}

func (this *priority_queue) Contains(i int) bool {
	return this.qp[i] != -1
}

func (this *priority_queue) ChangeKey(i int, key algorithms.Comparable) {
	if !this.Contains(i) {
		panic("index is not in the priority queue\n")
	}
	this.keys[i] = key
	this.swim(this.qp[i])
	this.sink(this.qp[i])
}

func (this *priority_queue) IncreaseKey(i int, key algorithms.Comparable) {
	if !this.Contains(i) {
		panic("index is not in the priority queue\n")
	}
	if this.keys[i].CompareTo(key) >= 0 {
		panic("Calling IncreaseKey() with given argument would not strictly increase the key\n")
	}

	this.keys[i] = key
	this.swim(this.qp[i])
}

func (this *priority_queue) DecreaseKey(i int, key algorithms.Comparable) {
	if !this.Contains(i) {
		panic("index is not in the priority queue\n")
	}
	if this.keys[i].CompareTo(key) <= 0 {
		panic("Calling DecreaseKey() with given argument would not strictly decrease the key\n")
	}

	this.keys[i] = key
	this.sink(this.qp[i])
}

func (this *priority_queue) Iterator() Iterator {
	this.current = 1
	return this
}

func (this *priority_queue) HasNext() bool {
	return this.current != this.size+1
}

func (this *priority_queue) Next() *Item {
	item := &Item{}
	item.Next = nil
	item.Value = &PriorityQueueItem{this.pq[this.current], this.keys[this.pq[this.current]]}
	this.current++
	return item
}

func (this *priority_queue) comp(i, j int) bool {
	if this.max_pq {
		return this.keys[this.pq[i]].CompareTo(this.keys[this.pq[j]]) < 0
	} else {
		return this.keys[this.pq[i]].CompareTo(this.keys[this.pq[j]]) > 0
	}
}

func (this *priority_queue) swap(i, j int) {
	t := this.pq[i]
	this.pq[i] = this.pq[j]
	this.pq[j] = t

	this.qp[this.pq[i]] = i
	this.qp[this.pq[j]] = j
}

func (this *priority_queue) swim(k int) {
	for k > 1 && this.comp(k/2, k) {
		this.swap(k/2, k)
		k = k / 2
	}
}

func (this *priority_queue) sink(k int) {
	for 2*k <= this.size {
		j := 2 * k
		if j < this.size && this.comp(j, j+1) {
			j++
		}
		if !this.comp(k, j) {
			break
		}
		this.swap(k, j)
		k = j
	}
}
