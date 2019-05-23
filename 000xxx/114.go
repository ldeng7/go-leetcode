func cal(n *TreeNode) *TreeNode {
	r := n.Right
	n.Right, n.Left = n.Left, nil
	t := n
	if nil != n.Right {
		t = cal(n.Right)
	}
	t.Right = r
	if nil != r {
		t = cal(r)
	}
	return t
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	cal(root)
}
