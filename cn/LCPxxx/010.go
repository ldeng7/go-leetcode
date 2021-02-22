func max(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}

func cal(n *TreeNode) (int, float64) {
	if nil == n {
		return 0, 0
	}
	il, fl := cal(n.Left)
	ir, fr := cal(n.Right)
	return n.Val + il + ir, float64(n.Val) + max(max(fl, fr), float64(il+ir)/2)
}

func minimalExecTime(root *TreeNode) float64 {
	_, o := cal(root)
	return o
}
