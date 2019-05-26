func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if nil == root || nil == root.Left {
		return root
	}
	l, r := root.Left, root.Right
	out := upsideDownBinaryTree(l)
	l.Left, l.Right = r, root
	root.Left, root.Right = nil, nil
	return out
}
