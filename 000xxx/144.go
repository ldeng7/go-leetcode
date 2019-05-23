func preorderTraversal(root *TreeNode) []int {
	n := root
	out := []int{}
	st := []*TreeNode{}
	for nil != n || 0 != len(st) {
		if nil != n {
			out = append(out, n.Val)
			st = append(st, n.Right)
			n = n.Left
		} else {
			n = st[len(st)-1]
			st = st[:len(st)-1]
		}
	}
	return out
}
