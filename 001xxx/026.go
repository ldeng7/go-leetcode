func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func maxAncestorDiff(root *TreeNode) int {
	o := 0
	var cal func(*TreeNode, int, int)
	cal = func(n *TreeNode, l, h int) {
		o = max(abs(n.Val-l), o)
		o = max(abs(n.Val-h), o)
		l = min(l, n.Val)
		h = max(h, n.Val)
		if nil != n.Left {
			cal(n.Left, l, h)
		}
		if nil != n.Right {
			cal(n.Right, l, h)
		}
	}
	cal(root, root.Val, root.Val)
	return o
}
