type FindElements struct {
	r *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	return FindElements{root}
}

func (this *FindElements) Find(target int) bool {
	target++
	p, m, n := this.r, 1, target
	for n != 1 {
		m, n = m<<1, n>>1
	}
	for m >>= 1; m != 0 && p != nil; m >>= 1 {
		if m&target != 0 {
			p = p.Right
		} else {
			p = p.Left
		}
	}
	return p != nil
}
