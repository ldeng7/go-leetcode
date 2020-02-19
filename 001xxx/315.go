func cal(n *TreeNode, p, g int) int {
	if nil != n {
		o := cal(n.Left, n.Val, p) + cal(n.Right, n.Val, p)
		if g&1 == 0 {
			o += n.Val
		}
		return o
	}
	return 0
}

func sumEvenGrandparent(root *TreeNode) int {
	return cal(root, 1, 1)
}
