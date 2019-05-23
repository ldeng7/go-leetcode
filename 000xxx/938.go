func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	var cal func(*TreeNode, int, int)
	cal = func(node *TreeNode, l, r int) {
		if nil == node {
			return
		}
		if node.Val < l {
			cal(node.Right, l, r)
		} else if node.Val > r {
			cal(node.Left, l, r)
		} else {
			sum += node.Val
			cal(node.Right, l, r)
			cal(node.Left, l, r)
		}
	}
	cal(root, L, R)
	return sum
}
