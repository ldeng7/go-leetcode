func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if nil == root1 && nil == root2 {
		return true
	} else if nil == root1 || nil == root2 || root1.Val != root2.Val {
		return false
	}
	l1, r1, l2, r2 := root1.Left, root1.Right, root2.Left, root2.Right
	return (flipEquiv(l1, l2) && flipEquiv(r1, r2)) || (flipEquiv(l1, r2) && flipEquiv(l2, r1))
}
