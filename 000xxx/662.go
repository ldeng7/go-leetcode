func widthOfBinaryTree(root *TreeNode) int {
	out, start := 0, []int{}
	var cal func(*TreeNode, int, int)
	cal = func(n *TreeNode, h, i int) {
		if nil == n {
			return
		}
		if h >= len(start) {
			start = append(start, i)
		}
		o := i - start[h] + 1
		if o > out {
			out = o
		}
		cal(n.Left, h+1, i<<1)
		cal(n.Right, h+1, i<<1+1)

	}
	cal(root, 0, 1)
	return out
}
