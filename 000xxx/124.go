import "math"

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func cal(n *TreeNode, o *int) int {
	if nil == n {
		return 0
	}
	l := max(cal(n.Left, o), 0)
	r := max(cal(n.Right, o), 0)
	*o = max(*o, l+r+n.Val)
	return max(l, r) + n.Val
}

func maxPathSum(root *TreeNode) int {
	o := math.MinInt64
	cal(root, &o)
	return o
}
