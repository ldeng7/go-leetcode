func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	b := false
	for 0 != len(q) {
		n := q[0]
		q = q[1:]
		if nil == n {
			b = true
			continue
		}
		if b {
			return false
		}
		q = append(q, n.Left, n.Right)
	}
	return true
}
