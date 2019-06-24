func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func distributeCoins(root *TreeNode) int {
	o := 0
	var cal func(*TreeNode) int
	cal = func(n *TreeNode) int {
		if nil == n {
			return 0
		}
		l, r := cal(n.Left), cal(n.Right)
		o += abs(l) + abs(r)
		return n.Val + l + r - 1
	}
	cal(root)
	return o
}
