func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func findLeaves(root *TreeNode) [][]int {
	out := [][]int{}
	var cal func(*TreeNode) int
	cal = func(n *TreeNode) int {
		if nil == n {
			return -1
		}
		d := max(cal(n.Left), cal(n.Right)) + 1
		for d >= len(out) {
			out = append(out, []int{})
		}
		out[d] = append(out[d], n.Val)
		return d
	}
	cal(root)
	return out
}
