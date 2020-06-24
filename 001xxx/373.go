func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxSumBST(root *TreeNode) int {
	o := 0
	var cal func(*TreeNode) (int, int, int, bool)
	cal = func(n *TreeNode) (int, int, int, bool) {
		if nil == n {
			return 0, 40001, -40001, true
		}
		ls, ll, lr, lb := cal(n.Left)
		rs, rl, rr, rb := cal(n.Right)
		if v := n.Val; lb && rb && v > lr && v < rl {
			s := v + ls + rs
			o = max(o, s)
			return s, min(ll, v), max(rr, v), true
		}
		return 0, 0, 0, false
	}
	cal(root)
	return o
}
