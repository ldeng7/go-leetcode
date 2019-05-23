func inorderTraversal(root *TreeNode) []int {
	n := root
	out := []int{}
	st := []*TreeNode{}
	for nil != n || 0 != len(st) {
		if nil != n {
			st = append(st, n)
			n = n.Left
		} else {
			n = st[len(st)-1]
			st = st[:len(st)-1]
			out = append(out, n.Val)
			n = n.Right
		}
	}
	return out
}
