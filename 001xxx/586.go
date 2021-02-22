type BSTIterator struct {
	ar []int
	i  int
}

func (this *BSTIterator) cal(n *TreeNode) {
	if nil == n {
		return
	}
	this.cal(n.Left)
	this.ar = append(this.ar, n.Val)
	this.cal(n.Right)
}

func Constructor(root *TreeNode) BSTIterator {
	this := &BSTIterator{make([]int, 0, 16), -1}
	this.cal(root)
	return *this
}

func (this *BSTIterator) HasNext() bool {
	return this.i < len(this.ar)-1
}

func (this *BSTIterator) Next() int {
	this.i++
	return this.ar[this.i]
}

func (this *BSTIterator) HasPrev() bool {
	return this.i > 0
}

func (this *BSTIterator) Prev() int {
	this.i--
	return this.ar[this.i]
}
