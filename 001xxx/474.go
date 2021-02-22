func deleteNodes(head *ListNode, m int, n int) *ListNode {
	node, c := head, 1
	for nil != node {
		for c < m && nil != node.Next {
			c++
			node = node.Next
		}
		for c < m+n && nil != node.Next {
			c++
			node.Next = node.Next.Next
		}
		c = 1
		node = node.Next
	}
	return head
}
