func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if nil == t1 {
		return t2
	} else if nil == t2 {
		return t1
	}
	return &TreeNode{t1.Val + t2.Val, mergeTrees(t1.Left, t2.Left), mergeTrees(t1.Right, t2.Right)}
}
