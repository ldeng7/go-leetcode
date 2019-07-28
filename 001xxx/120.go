func maximumAverageSubtree(root *TreeNode) float64 {
	var o float64
	var cal func(*TreeNode) (int, int)
	cal = func(node *TreeNode) (int, int) {
		if nil == node {
			return 0, 0
		}
		sl, nl := cal(node.Left)
		sr, nr := cal(node.Right)
		s, n := node.Val+sl+sr, 1+nl+nr
		if a := float64(s) / float64(n); a > o {
			o = a
		}
		return s, n
	}
	cal(root)
	return o
}
