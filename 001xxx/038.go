func cal(n *TreeNode, sum *int) {
	if nil == n {
		return
	}
	cal(n.Right, sum)
	*sum += n.Val
	n.Val = *sum
	cal(n.Left, sum)
}
func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	cal(root, &sum)
	return root
}
