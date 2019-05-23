func invertTree(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
