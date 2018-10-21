func minDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	dl := minDepth(root.Left)
	dr := minDepth(root.Right)
	if dl == 0 {
		return 1 + dr
	}
	if dr == 0 {
		return 1 + dl
	}
	d := dl
	if dr < dl {
		d = dr
	}
	return 1 + d
}
