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

func reverseKGroup(head *ListNode, k int) *ListNode {
	l := head
	for i := 0; i < k-1; i++ {
		if nil == l {
			return head
		}
		l = l.Next
	}
	if nil == l {
		return head
	}
	n := l.Next
	l.Next = nil
	reverseList(head)
	head.Next = reverseKGroup(n, k)
	return l
}
