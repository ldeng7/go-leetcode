func do(root *TreeNode, h *TreeNode) *TreeNode {
	if nil == root {
		return h
	}
	left, right := root.Left, root.Right
	h.Right = root
	h = h.Right
	h.Left, h.Right = nil, nil
	return do(right, do(left, h))
}

func flatten(root *TreeNode) {
	do(root, &TreeNode{})
}
