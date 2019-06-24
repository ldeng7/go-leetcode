func findBottomLeftValue(root *TreeNode) int {
	q := []*TreeNode{root}
	var n *TreeNode
	for 0 != len(q) {
		n, q = q[0], q[1:]
		if nil != n.Right {
			q = append(q, n.Right)
		}
		if nil != n.Left {
			q = append(q, n.Left)
		}
	}
	return n.Val
}
