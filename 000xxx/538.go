func convertBST(root *TreeNode) *TreeNode {
	s := 0
	var cal func(*TreeNode)
	cal = func(n *TreeNode) {
		if nil == n {
			return
		}
		cal(n.Right)
		n.Val += s
		s = n.Val
		cal(n.Left)
	}
	cal(root)
	return root
}
