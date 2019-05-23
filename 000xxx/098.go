func check(node *TreeNode, min, max *int) bool {
	if nil == node {
		return true
	}
	if (nil != min && node.Val <= *min) || (nil != max && node.Val >= *max) {
		return false
	}
	return check(node.Left, min, &node.Val) && check(node.Right, &node.Val, max)
}

func isValidBST(root *TreeNode) bool {
	if nil == root {
		return true
	}
	return check(root.Left, nil, &root.Val) && check(root.Right, &root.Val, nil)
}
