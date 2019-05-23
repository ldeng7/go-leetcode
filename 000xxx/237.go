func deleteNode(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
}
