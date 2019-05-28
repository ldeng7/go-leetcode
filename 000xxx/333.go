import "math"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func cal(n *TreeNode) (int, int, int) {
	if nil == n {
		return math.MaxInt64, math.MinInt64, 0
	}
	lmin, lmax, ln := cal(n.Left)
	rmin, rmax, rn := cal(n.Right)
	if n.Val > lmax && n.Val < rmin {
		return min(n.Val, lmin), max(n.Val, rmax), ln + rn + 1
	}
	return math.MinInt64, math.MaxInt64, max(ln, rn)
}

func largestBSTSubtree(root *TreeNode) int {
	_, _, n := cal(root)
	return n
}
