func sortedArrayToBST(nums []int) *TreeNode {
	if 0 == len(nums) {
		return nil
	}
	i := len(nums) >> 1
	node := &TreeNode{Val: nums[i]}
	node.Left = sortedArrayToBST(nums[:i])
	node.Right = sortedArrayToBST(nums[i+1:])
	return node
}

func sortedListToBST(head *ListNode) *TreeNode {
	nums := []int{}
	for nil != head {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return sortedArrayToBST(nums)
}
