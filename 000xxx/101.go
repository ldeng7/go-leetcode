func cal(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if (l != nil && r == nil) || (l == nil && r != nil) {
		return false
	}
	if l.Val != r.Val {
		return false
	}
	if !cal(l.Left, r.Right) {
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
