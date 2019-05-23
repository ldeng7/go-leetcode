func buildTree(inorder []int, postorder []int) *TreeNode {
	var cal func(il, ir, pl, pr int) *TreeNode
	cal = func(il, ir, pl, pr int) *TreeNode {
		if il > ir || pl > pr {
			return nil
		}
		i := il
		for ; i < ir; i++ {
			if inorder[i] == postorder[pr] {
				break
			}
		}
		node := &TreeNode{Val: postorder[pr]}
		node.Left = cal(il, i-1, pl, pl+i-il-1)
		node.Right = cal(i+1, ir, pl+i-il, pr-1)
		return node
	}
	return cal(0, len(inorder)-1, 0, len(postorder)-1)
}
