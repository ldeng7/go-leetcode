func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func cal(n *TreeNode) (int, *TreeNode) {
	if nil == n {
		return 0, nil
	}
	d1, n1 := cal(n.Left)
	d2, n2 := cal(n.Right)
	d := max(d1, d2) + 1
	if d1 > d2 {
		n = n1
	} else if d2 > d1 {
		n = n2
	}
	return d, n
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	_, n := cal(root)
	return n
}
