type CBTInserter struct {
	r, ins *TreeNode
	q      []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	ins, q := root, []*TreeNode{root}
	for 0 != len(q) {
		n := q[0]
		q = q[1:]
		if nil == n.Left {
			ins = n
			break
		}
		if nil == n.Right {
			q = append(q, n.Left)
			ins = n
			break
		}
		q = append(q, n.Left, n.Right)
	}
	return CBTInserter{root, ins, q}
}

func (this *CBTInserter) Insert(v int) int {
	o := this.ins.Val
	n := &TreeNode{Val: v}
	if nil == this.ins.Left {
		this.ins.Left, this.q = n, append(this.q, n)
	} else {
		this.ins.Right = n
		this.ins = this.q[0]
		this.q = append(this.q[1:], n)
	}
	return o
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.r
}
