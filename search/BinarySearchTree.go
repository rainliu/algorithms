package search

import (
	"algorithms"
	"algorithms/container"
)

type Node struct{
	key	algorithms.Comparable
	val interface{}
	left	*Node
	right	*Node
	size	int
}

func NewNode(key algorithms.Comparable, val interface{}, size int) *Node{
	this := &Node{}
	this.key = key;
	this.val = val;
	this.size = size;
	return this
}

type BinarySearchTree struct{
	root *Node
}

func (this *BinarySearchTree) IsEmpty() bool{
	return this.Size()==0
}

func (this *BinarySearchTree) Contains(key algorithms.Comparable) bool{
	return this.Get(key) != nil
}

func (this *BinarySearchTree) RangedSize(low, high algorithms.Comparable) int{
	if high.CompareTo(low) < 0 {
		return 0
	}else if this.Contains(high) {
		return this.Rank(high) - this.Rank(low) + 1
	}else{
		return this.Rank(high) - this.Rank(low)
	}
}

func (this *BinarySearchTree) Size() int{
	return this.size(this.root)
}

func (this *BinarySearchTree) size(x *Node) int{
	if x==nil {
		return 0
	}else{
		return x.size
	}
}

func (this *BinarySearchTree) Get(key algorithms.Comparable) interface{}{
	return this.get(this.root, key)
}

func (this *BinarySearchTree) get(x *Node, key algorithms.Comparable) interface{}{
	if x==nil{
		return nil;
	}

	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		return this.get(x.left, key)
	}else if cmp > 0 {
		return this.get(x.right, key)
	}else{
		return x.val;
	}
}

func (this *BinarySearchTree) Set(key algorithms.Comparable, val interface{}){
	this.root = this.set(this.root, key, val)
}

func (this *BinarySearchTree) set(x *Node, key algorithms.Comparable, val interface{}) *Node{
	if x==nil{
		return NewNode(key, val, 1);
	}
	
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		x.left = this.set(x.left, key, val)
	}else if cmp > 0 {
		x.right= this.set(x.right, key, val)
	}else{
		x.val  = val; 
	}
	
	x.size = this.size(x.left) + this.size(x.right) + 1;
	
	return x;
}

func (this *BinarySearchTree) Min() algorithms.Comparable {
	return this.min(this.root).key
}

func (this *BinarySearchTree) min(x *Node) *Node{
	if x.left == nil {
		return x
	}else{
		return this.min(x.left)
	}
}

func (this *BinarySearchTree) Floor(key algorithms.Comparable) algorithms.Comparable{
	x := this.floor(this.root, key)
	if x==nil {
		return nil
	}
	
	return x.key
}

func (this *BinarySearchTree) floor(x *Node, key algorithms.Comparable) *Node{
	if x==nil{
		return nil
	}
	
	cmp := key.CompareTo(x.key)
	if cmp == 0 {
		return x
	}else if cmp < 0 {
		return this.floor(x.left, key)
	}else{	
		t:= this.floor(x.right, key)
		if t!=nil{
			return t
		}else{
			return x
		}
	}
}

func (this *BinarySearchTree) Max() algorithms.Comparable {
	return this.max(this.root).key
}

func (this *BinarySearchTree) max(x *Node) *Node{
	if x.right == nil {
		return x
	}else{
		return this.max(x.right)
	}
}


func (this *BinarySearchTree) Ceil(key algorithms.Comparable) algorithms.Comparable{
	x := this.ceil(this.root, key)
	if x==nil {
		return nil
	}
	
	return x.key
}

func (this *BinarySearchTree) ceil(x *Node, key algorithms.Comparable) *Node{
	if x==nil{
		return nil
	}
	
	cmp := key.CompareTo(x.key)
	if cmp == 0 {
		return x
	}else if cmp > 0 {
		return this.ceil(x.right, key)
	}else{	
		t:= this.ceil(x.left, key)
		if t!=nil{
			return t
		}else{
			return x
		}
	}
}

func (this *BinarySearchTree) Select(k int) algorithms.Comparable{ 
	return this._select(this.root, k).key
}

func (this *BinarySearchTree) _select(x *Node, k int) *Node{
	if x==nil{
		return nil
	}
	t := this.size(x.left)
	if t>k {
		return this._select(x.left, k)
	}else if t < k {
		return this._select(x.right, k-t-1)
	}else{
		return x;
	}
}

func (this *BinarySearchTree) Rank(key algorithms.Comparable) int{
	return this.rank(this.root, key)
}

func (this *BinarySearchTree) rank(x *Node, key algorithms.Comparable) int{
	if x==nil {
		return 0
	}
	
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		return this.rank(x.left, key)
	}else if cmp > 0 {
		return this.rank(x.right, key) + this.size(x.left) + 1
	}else {
		return this.size(x.left)
	}
}

func (this *BinarySearchTree) DeleteMin(){
	this.root = this.deleteMin(this.root)
}

func (this *BinarySearchTree) deleteMin(x *Node) *Node{
	if x.left == nil{
		return x.right
	}
	x.left = this.deleteMin(x.left)
	x.size = this.size(x.left) + this.size(x.right) + 1
	return x
}

func (this *BinarySearchTree) DeleteMax(){
	this.root = this.deleteMax(this.root)
}

func (this *BinarySearchTree) deleteMax(x *Node) *Node{
	if x.right == nil{
		return x.left
	}
	x.right = this.deleteMax(x.right)
	x.size = this.size(x.left) + this.size(x.right) + 1
	return x
}

func (this *BinarySearchTree) Delete(key algorithms.Comparable){
	this.root = this._delete(this.root, key)
}

func (this *BinarySearchTree) _delete(x *Node, key algorithms.Comparable) *Node{
	if x==nil {
		return nil
	}
	cmp := key.CompareTo(x.key)
	if cmp < 0 {
		x.left = this._delete(x.left, key)
	}else if cmp > 0 {
		x.right = this._delete(x.right, key)
	}else{
		if x.right == nil{
			return x.left
		}
		if x.left == nil{
			return x.right
		}
		t := x;
		x = this.min(t.right)
		x.right = this.deleteMin(t.right)
		x.left = t.left
	}
	
	x.size = this.size(x.left) + this.size(x.right) +1
	
	return x
}

func (this *BinarySearchTree) Keys() container.Iterator{
	return this.RangedKeys(this.Min(), this.Max())
}

func (this *BinarySearchTree) RangedKeys(low, high algorithms.Comparable) container.Iterator{
	queue := &container.Queue{}
	
	this.keys(this.root, queue, low, high);
	
	return queue.Iterator();
}

func (this *BinarySearchTree) keys(x *Node, queue *container.Queue, low, high algorithms.Comparable){
	if x==nil {
		return
	}
	cmplow := low.CompareTo(x.key)
	cmphigh:= high.CompareTo(x.key)
	if cmplow < 0 {
		this.keys(x.left, queue, low, high)
	}
	if cmplow <=0 && cmphigh >=0 {
		queue.Push(x.key)
	}
	if cmphigh > 0 {
		this.keys(x.right, queue, low, high)
	}
}