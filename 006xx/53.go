func do(node *TreeNode, k int, m map[int]bool) bool {
	if nil == node {
		return false
	}
	if _, ok := m[k-node.Val]; ok {
		return true
	}
	m[node.Val] = true
	return do(node.Left, k, m) || do(node.Right, k, m)
}

func findTarget(root *TreeNode, k int) bool {
	m := map[int]bool{}
	return do(root, k, m)
}
