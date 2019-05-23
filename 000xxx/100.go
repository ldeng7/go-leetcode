func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	}
	if (nil == p && nil != q) || (nil != p && nil == q) || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
