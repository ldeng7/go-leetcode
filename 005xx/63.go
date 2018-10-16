func sum(node *TreeNode) int {
	if nil == node {
		return 0
	}
	return node.Val + sum(node.Left) + sum(node.Right)
}

func findTilt(root *TreeNode) int {
	if nil == root {
		return 0
	}
	l, r := 0, 0
	if nil != root.Left {
		l = sum(root.Left)
	}
	if nil != root.Right {
		r = sum(root.Right)
	}
	diff := l - r
	if diff < 0 {
		diff = -diff
	}
	return diff + findTilt(root.Left) + findTilt(root.Right)
}
