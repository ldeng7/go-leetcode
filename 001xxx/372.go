func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func longestZigZag(root *TreeNode) int {
	o := 0
	var cal func(*TreeNode) (int, int)
	cal = func(n *TreeNode) (int, int) {
		r, l := 0, 0
		if nil != n.Left {
			_, l = cal(n.Left)
			l++
		}
		if nil != n.Right {
			r, _ = cal(n.Right)
			r++
		}
		o = max(o, max(l, r))
		return l, r
	}
	cal(root)
	return o
}
