func bstFromPreorder(preorder []int) *TreeNode {
	if 0 == len(preorder) {
		return nil
	}
	i := 1
	for ; i < len(preorder); i++ {
		if preorder[i] > preorder[0] {
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:i]),
		Right: bstFromPreorder(preorder[i:]),
	}
}
