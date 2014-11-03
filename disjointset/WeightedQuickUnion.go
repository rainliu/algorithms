package disjointset

import ()

type WeightedQuickUnion struct {
	disjointset
	sz []int
}

func NewWeightedQuickUnion(N int) *WeightedQuickUnion {
	this := &WeightedQuickUnion{}
	this.UnionFind = this
	this.Init(N)
	this.sz = make([]int, N)
	for i := 0; i < N; i++ {
		this.sz[i] = 1
	}
	return this
}

func (this *WeightedQuickUnion) Find(p int) int {
	for p != this.id[p] {
		p = this.id[p]
	}
	return p
}

func (this *WeightedQuickUnion) Union(p, q int) {
	pRoot := this.Find(p)
	qRoot := this.Find(q)

	if pRoot == qRoot {
		return
	}

	if this.sz[pRoot] < this.sz[qRoot] {
		this.id[pRoot] = qRoot
		this.sz[qRoot] += this.sz[pRoot]
	} else {
		this.id[qRoot] = pRoot
		this.sz[pRoot] += this.sz[qRoot]
	}

	this.count--
}
