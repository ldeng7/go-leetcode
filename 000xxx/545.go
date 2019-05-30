func boundaryOfBinaryTree(root *TreeNode) []int {
	if nil == root {
		return []int{}
	}
	out := []int{root.Val}

	var cal func(n *TreeNode, l, r bool)
	cal = func(n *TreeNode, l, r bool) {
		if nil == n {
			return
		}
		if nil == n.Left && nil == n.Right {
			out = append(out, n.Val)
			return
		}
		if l {
			out = append(out, n.Val)
		}
		cal(n.Left, l && nil != n.Left, r && nil == n.Right)
		cal(n.Right, l && nil == n.Left, r && nil != n.Right)
		if r {
			out = append(out, n.Val)
		}
	}

	cal(root.Left, true, false)
	cal(root.Right, false, true)
	return out
}
