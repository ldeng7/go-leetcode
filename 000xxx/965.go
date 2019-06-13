func isUnivalTree(root *TreeNode) bool {
	v := root.Val
	var cal func(n *TreeNode) bool
	cal = func(n *TreeNode) bool {
		return n == nil || (n.Val == v && cal(n.Left) && cal(n.Right))
	}
	return cal(root)
}
