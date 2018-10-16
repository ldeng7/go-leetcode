func searchBST(root *TreeNode, val int) *TreeNode {
	if nil == root {
		return nil
	}
	if root.Val == val {
		return root
	}
	node := searchBST(root.Left, val)
	if nil != node {
		return node
	}
	return searchBST(root.Right, val)
}
