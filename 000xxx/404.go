func cal(node *TreeNode, isLeft bool, sum *int) {
	if nil == node {
		return
	}
	if isLeft && nil == node.Left && nil == node.Right {
		*sum += node.Val
	}
	cal(node.Left, true, sum)
	cal(node.Right, false, sum)
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	cal(root, false, &sum)
	return sum
}
