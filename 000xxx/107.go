func levelOrderBottom(root *TreeNode) [][]int {
	out := [][]int{}
	var cal func(*TreeNode, int)
	cal = func(node *TreeNode, level int) {
		if nil == node {
			return
		}
		if level == len(out) {
			out = append(out, []int{})
		}
		out[level] = append(out[level], node.Val)
		cal(node.Left, level+1)
		cal(node.Right, level+1)
	}
	cal(root, 0)
	out1 := make([][]int, len(out))
	for i, ar := range out {
		out1[len(out)-i-1] = ar
	}
	return out1
}
