import "math"

func getMinimumDifference(root *TreeNode) int {
	o, pv := math.MaxInt64, -1
	var cal func(*TreeNode)
	cal = func(n *TreeNode) {
		if nil == n {
			return
		}
		cal(n.Left)
		if (pv != -1) && n.Val-pv < o {
			o = n.Val - pv
		}
		pv = n.Val
		cal(n.Right)
	}
	cal(root)
	return o
}
