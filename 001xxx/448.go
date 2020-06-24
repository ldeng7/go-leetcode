func goodNodes(root *TreeNode) int {
	o := 0
	var cal func(*TreeNode, int)
	cal = func(n *TreeNode, m int) {
		if nil == n {
			return
		}
		if n.Val >= m {
			o++
			m = n.Val
		}
		cal(n.Left, m)
		cal(n.Right, m)
	}
	cal(root, -10000)
	return o
}
