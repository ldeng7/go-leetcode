func constructMaximumBinaryTree(nums []int) *TreeNode {
	ar := []*TreeNode{}
	for _, n := range nums {
		node := &TreeNode{Val: n}
		for 0 != len(ar) && ar[len(ar)-1].Val < n {
			node.Left = ar[len(ar)-1]
			ar = ar[:len(ar)-1]
		}
		if 0 != len(ar) {
			ar[len(ar)-1].Right = node
		}
		ar = append(ar, node)
	}
	return ar[0]
}
