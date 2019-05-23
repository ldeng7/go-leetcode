func pathSum(root *TreeNode, sum int) int {
	if nil == root {
		return 0
	}
	return cal(root, 0, sum) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func cal(n *TreeNode, pre, sum int) int {
	if nil == n {
		return 0
	}
	cur := pre + n.Val
	out := 0
	if cur == sum {
		out++
	}
	return out + cal(n.Left, cur, sum) + cal(n.Right, cur, sum)
}
