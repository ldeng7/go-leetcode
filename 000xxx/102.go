func levelOrder(root *TreeNode) [][]int {
	out := [][]int{}
	var cal func(*TreeNode, int)
	cal = func(node *TreeNode, level int) {
		if nil == node {
			return
		}
		if level == len(out) {
			out = append(out, nil)
		}
		out[level] = append(out[level], node.Val)
		cal(node.Left, level+1)
		cal(node.Right, level+1)
	}
	cal(root, 0)
	return out
}
