func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	l, r := 0, 0
	var cal func(*TreeNode) int
	cal = func(node *TreeNode) int {
		if nil == node {
			return 0
		}
		l1, r1 := cal(node.Left), cal(node.Right)
		if x == node.Val {
			l, r = l1, r1
		}
		return l1 + r1 + 1
	}

	s := cal(root)
	return s > (l+r+1)<<1 || s < l<<1 || s < r<<1
}
