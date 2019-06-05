func countNodes(root *TreeNode) int {
	var hl, hr uint
	l, r := root, root
	for nil != l {
		hl++
		l = l.Left
	}
	for nil != r {
		hr++
		r = r.Right
	}
	if hl == hr {
		return (1 << hl) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
