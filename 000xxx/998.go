func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	n := &TreeNode{Val: val}
	if nil == root || root.Val < val {
		n.Left = root
		return n
	}
	t := root
	for nil != t.Right && t.Right.Val > val {
		t = t.Right
	}
	n.Left, t.Right = t.Right, n
	return root
}
