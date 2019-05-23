func postorderTraversal(root *TreeNode) []int {
	n := root
	out := []int{}
	st := []*TreeNode{}
	for nil != n || 0 != len(st) {
		if nil != n {
			out = append(out, n.Val)
			st = append(st, n.Left)
			n = n.Right
		} else {
			n = st[len(st)-1]
			st = st[:len(st)-1]
		}
	}

	l := len(out)
	out1 := make([]int, l)
	for i, e := range out {
		out1[l-i-1] = e
	}
	return out1
}
