func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func longestUnivaluePath(root *TreeNode) int {
	if nil == root {
		return 0
	}
	var cal func(*TreeNode, int) int
	cal = func(n *TreeNode, pv int) int {
		if nil == n || n.Val != pv {
			return 0
		}
		return max(cal(n.Left, n.Val), cal(n.Right, n.Val)) + 1
	}
	return max(max(longestUnivaluePath(root.Left), longestUnivaluePath(root.Right)),
		cal(root.Left, root.Val)+cal(root.Right, root.Val))
}
