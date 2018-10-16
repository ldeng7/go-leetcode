func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	dl := maxDepth(root.Left)
	dr := maxDepth(root.Right)
	d := dl
	if dr > dl {
		d = dr
	}
	return 1 + d
}
