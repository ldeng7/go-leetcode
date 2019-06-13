func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := &TreeNode{Val: pre[0]}
	if len(pre) == 1 {
		return root
	}
	i, v1 := 0, pre[1]
	for j, v := range post {
		if v == v1 {
			i = j
			break
		}
	}
	root.Left = constructFromPrePost(pre[1:i+2], post[:i+1])
	root.Right = constructFromPrePost(pre[i+2:], post[i+1:len(post)-1])
	return root
}
