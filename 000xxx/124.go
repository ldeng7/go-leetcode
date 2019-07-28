import "math"

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxPathSum(root *TreeNode) int {
	o := math.MinInt64
	var cal func(*TreeNode) int
	cal = func(n *TreeNode) int {
		if nil == n {
			return 0
		}
		l := max(cal(n.Left), 0)
		r := max(cal(n.Right), 0)
		o = max(o, l+r+n.Val)
		return max(l, r) + n.Val
	}
	cal(root)
	return o
}
