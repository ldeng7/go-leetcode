import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minCameraCover(root *TreeNode) int {
	var cal func(*TreeNode) (int, int, int)
	cal = func(n *TreeNode) (int, int, int) {
		if nil == n {
			return 0, 0, math.MaxInt32
		}
		l0, l1, l2 := cal(n.Left)
		r0, r1, r2 := cal(n.Right)
		ml, mr := min(l1, l2), min(r1, r2)
		o0 := l1 + r1
		o1 := min(l2+mr, r2+ml)
		o2 := 1 + min(l0, ml) + min(r0, mr)
		return o0, o1, o2
	}
	_, o1, o2 := cal(root)
	return min(o1, o2)
}
