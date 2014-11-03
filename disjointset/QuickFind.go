package disjointset

import ()

type QuickFind struct {
	disjointset
}

func NewQuickFind(N int) *QuickFind {
	this := &QuickFind{}
	this.UnionFind = this
	this.Init(N)
	return this
}

func (this *QuickFind) Find(p int) int {
	return this.id[p]
}

func (this *QuickFind) Union(p, q int) {
	pID := this.Find(p)
	qID := this.Find(q)

	if pID == qID {
		return
	}

	for i := 0; i < len(this.id); i++ {
		if this.id[i] == pID {
			this.id[i] = qID
		}
	}

	this.count--
}
