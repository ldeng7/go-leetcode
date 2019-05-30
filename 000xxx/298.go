func cal(n *TreeNode, v, c int, out *int) {
	if nil == n {
		return
	}
	if n.Val == v+1 {
		c++
	} else {
		c = 1
	}
	if c > *out {
		*out = c
	}
	cal(n.Left, n.Val, c, out)
	cal(n.Right, n.Val, c, out)
}

func longestConsecutive(root *TreeNode) int {
	if nil == root {
		return 0
	}
	out := 0
	cal(root, root.Val, 0, &out)
	return out
}
