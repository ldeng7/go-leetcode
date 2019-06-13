func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func findSecondMinimumValue(root *TreeNode) int {
	if nil == root.Left {
		return -1
	}
	var l, r int
	if root.Left.Val == root.Val {
		l = findSecondMinimumValue(root.Left)
	} else {
		l = root.Left.Val
	}
	if root.Right.Val == root.Val {
		r = findSecondMinimumValue(root.Right)
	} else {
		r = root.Right.Val
	}
	if l != -1 && r != -1 {
		return min(l, r)
	}
	return max(l, r)
}
