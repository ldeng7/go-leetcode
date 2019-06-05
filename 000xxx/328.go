func oddEvenList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	p, n := head, head.Next
	for nil != n && nil != n.Next {
		t := p.Next
		p.Next = n.Next
		n.Next = n.Next.Next
		p.Next.Next = t
		p, n = p.Next, n.Next
	}
	return head
}
