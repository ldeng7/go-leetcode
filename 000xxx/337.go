func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func cal(n *TreeNode, l, r *int) int {
	if nil == n {
		return 0
	}
	var ll, lr, rl, rr int
	*l = cal(n.Left, &ll, &lr)
	*r = cal(n.Right, &rl, &rr)
	return max(n.Val+ll+lr+rl+rr, *l+*r)
}

func rob(root *TreeNode) int {
	var l, r int
	return cal(root, &l, &r)
}
