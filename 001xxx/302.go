func deepestLeavesSum(root *TreeNode) int {
	o, lm := 0, 0
	var cal func(*TreeNode, int)
	cal = func(n *TreeNode, l int) {
		if nil == n {
			return
		}
		if l > lm {
			lm = l
			o = n.Val
		} else if l == lm {
			o += n.Val
		}
		cal(n.Left, l+1)
		cal(n.Right, l+1)
	}
	cal(root, 0)
	return o
}
