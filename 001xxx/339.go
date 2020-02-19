func maxProduct(root *TreeNode) int {
	m := map[*TreeNode]int{}
	var sum func(*TreeNode) int
	sum = func(n *TreeNode) int {
		if nil == n {
			return 0
		}
		s := n.Val + sum(n.Left) + sum(n.Right)
		m[n] = s
		return s
	}
	s := sum(root)

	o := 0
	var cal func(*TreeNode)
	cal = func(n *TreeNode) {
		if nil == n {
			return
		}
		v := m[n]
		if t := v * (s - v); t > o {
			o = t
		}
		cal(n.Left)
		cal(n.Right)
	}
	cal(root)
	return o % 1000000007
}
