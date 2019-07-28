func longestConsecutive(root *TreeNode) int {
	if nil == root {
		return 0
	}
	out := 0
	var cal func(*TreeNode, int, int)
	cal = func(n *TreeNode, v, c int) {
		if nil == n {
			return
		}
		if n.Val == v+1 {
			c++
		} else {
			c = 1
		}
		if c > out {
			out = c
		}
		cal(n.Left, n.Val, c)
		cal(n.Right, n.Val, c)
	}
	cal(root, root.Val, 0)
	return out
}
