func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func cal(n *TreeNode, d int) int {
	if nil == n {
		return 0
	}
	l, r := 0, 0
	if nil != n.Left && n.Val-n.Left.Val == d {
		l = cal(n.Left, d) + 1
	}
	if nil != n.Right && n.Val-n.Right.Val == d {
		r = cal(n.Right, d) + 1
	}
	return max(l, r)
}

func longestConsecutive(root *TreeNode) int {
	if nil == root {
		return 0
	}
	return max(cal(root, 1)+cal(root, -1)+1,
		max(longestConsecutive(root.Left), longestConsecutive(root.Right)))
}
