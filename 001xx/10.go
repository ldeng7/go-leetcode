func getDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	dl := getDepth(root.Left)
	dr := getDepth(root.Right)
	d := dl
	if dr > dl {
		d = dr
	}
	return 1 + d
}

func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}
	dl := getDepth(root.Left)
	dr := getDepth(root.Right)
	diff := dl - dr
	if diff > 1 || diff < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
