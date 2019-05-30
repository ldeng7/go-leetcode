import "math"

func closestKValues(root *TreeNode, target float64, k int) []int {
	out := []int{}
	st := []*TreeNode{}
	n := root
	for nil != n || 0 != len(st) {
		for nil != n {
			st = append(st, n)
			n = n.Left
		}
		n = st[len(st)-1]
		st = st[:len(st)-1]
		if len(out) < k {
			out = append(out, n.Val)
		} else if math.Abs(float64(n.Val)-target) < math.Abs(float64(out[0])-target) {
			out = append(out, n.Val)
			out = out[1:]
		} else {
			break
		}
		n = n.Right
	}
	return out
}
