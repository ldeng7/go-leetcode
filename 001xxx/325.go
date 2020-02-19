func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if nil != root.Left {
		root.Left = removeLeafNodes(root.Left, target)
	}
	if nil != root.Right {
		root.Right = removeLeafNodes(root.Right, target)
	}
	if root.Left == root.Right && root.Val == target {
		return nil
	}
	return root
}
