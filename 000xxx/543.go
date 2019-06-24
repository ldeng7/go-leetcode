func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func diameterOfBinaryTree(root *TreeNode) int {
	o, m := 0, map[*TreeNode]int{}
	var cal func(*TreeNode) int
	cal = func(n *TreeNode) int {
		if nil == n {
			return 0
		} else if d, ok := m[n]; ok {
			return d
		}
		l, r := cal(n.Left), cal(n.Right)
		o = max(o, l+r)
		m[n] = max(l, r) + 1
		return m[n]
	}
	cal(root)
	return o
}
