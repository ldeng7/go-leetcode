func checkEqualTree(root *TreeNode) bool {
	m := map[int]int{}
	var cal func(*TreeNode) int
	cal = func(n *TreeNode) int {
		if nil == n {
			return 0
		}
		v := n.Val + cal(n.Left) + cal(n.Right)
		m[v]++
		return v
	}
	s := cal(root)
	if 0 == s {
		return m[0] > 1
	}
	return s&1 == 0 && m[s>>1] != 0
}
