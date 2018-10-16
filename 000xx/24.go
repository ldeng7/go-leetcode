func swapPairs(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	h := head.Next
	head.Next = swapPairs(h.Next)
	h.Next = head
	return h
}
