func cal(root *TreeNode, h *TreeNode) *TreeNode {
	if nil == root {
		return h
	}
	left, right := root.Left, root.Right
	h.Right = root
	h = h.Right
	h.Left, h.Right = nil, nil
	return cal(right, cal(left, h))
}

func flatten(root *TreeNode) {
	cal(root, &TreeNode{})
}
