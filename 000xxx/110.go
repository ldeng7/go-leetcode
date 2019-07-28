func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func cal(n *TreeNode) (int, bool) {
	if nil == n {
		return 0, true
	}
	dl, bl := cal(n.Left)
	dr, br := cal(n.Right)
	if !bl || !br || abs(dl-dr) > 1 {
		return 0, false
	}
	return max(dl, dr) + 1, true
}

func isBalanced(root *TreeNode) bool {
	_, b := cal(root)
	return b
}
