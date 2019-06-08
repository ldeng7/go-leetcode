func cal(n *TreeNode) int {
	if nil == n {
		return 0
	}
	return 1 + cal(n.Left) + cal(n.Right)
}

func kthSmallest(root *TreeNode, k int) int {
	c := cal(root.Left)
	if k <= c {
		return kthSmallest(root.Left, k)
	} else if k > c+1 {
		return kthSmallest(root.Right, k-c-1)
	}
	return root.Val
}
