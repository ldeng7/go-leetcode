func cal(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if (l != nil && r == nil) || (l == nil && r != nil) ||
		l.Val != r.Val || !cal(l.Left, r.Right) {
		return false
	}
	return cal(l.Right, r.Left)
}

func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}
	return cal(root.Left, root.Right)
}
