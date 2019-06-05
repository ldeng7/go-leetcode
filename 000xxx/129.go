func cal(n *TreeNode, s int) int {
	if nil == n {
		return 0
	}
	s = s*10 + n.Val
	if nil == n.Left && nil == n.Right {
		return s
	}
	return cal(n.Left, s) + cal(n.Right, s)
}

func sumNumbers(root *TreeNode) int {
	return cal(root, 0)
}
