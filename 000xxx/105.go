func buildTree(preorder []int, inorder []int) *TreeNode {
	var cal func(pl, pr, il, ir int) *TreeNode
	cal = func(pl, pr, il, ir int) *TreeNode {
		if pl > pr || il > ir {
			return nil
		}
		i := il
		for ; i < ir; i++ {
			if preorder[pl] == inorder[i] {
				break
			}
		}
		node := &TreeNode{Val: preorder[pl]}
		node.Left = cal(pl+1, pl+i-il, il, i-1)
		node.Right = cal(pl+i-il+1, pr, i+1, ir)
		return node
	}
	return cal(0, len(preorder)-1, 0, len(inorder)-1)
}
