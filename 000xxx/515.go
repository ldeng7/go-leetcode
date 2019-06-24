func largestValues(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}
	out := []int{}
	var cal func(*TreeNode, int)
	cal = func(n *TreeNode, d int) {
		if len(out) < d {
			out = append(out, n.Val)
		} else if n.Val > out[d-1] {
			out[d-1] = n.Val
		}
		if nil != n.Left {
			cal(n.Left, d+1)
		}
		if nil != n.Right {
			cal(n.Right, d+1)
		}
	}
	cal(root, 1)
	return out
}
