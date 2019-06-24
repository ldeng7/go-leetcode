func increasingBST(root *TreeNode) *TreeNode {
	h := &TreeNode{}
	c := h
	var cal func(*TreeNode)
	cal = func(n *TreeNode) {
		if nil == n {
			return
		}
		cal(n.Left)
		c.Right, n.Left = n, nil
		c = n
		cal(n.Right)
	}
	cal(root)
	return h.Right
}
