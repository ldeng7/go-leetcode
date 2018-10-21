func reverseList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	l := head
	head = reverseList(l.Next)
	l.Next.Next = l
	l.Next = nil
	return head
}
