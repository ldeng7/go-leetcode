type BSTIterator struct {
	st []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	st := []*TreeNode{}
	for nil != root {
		st = append(st, root)
		root = root.Left
	}
	return BSTIterator{st}
}

func (this *BSTIterator) Next() int {
	l := len(this.st)
	n := this.st[l-1]
	this.st = this.st[:l-1]
	out := n.Val
	if nil != n.Right {
		n = n.Right
		for nil != n {
			this.st = append(this.st, n)
			n = n.Left
		}
	}
	return out
}

func (this *BSTIterator) HasNext() bool {
	return 0 != len(this.st)
}
