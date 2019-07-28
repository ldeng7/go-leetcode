func merge(l1, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	} else if nil == l2 {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = merge(l1.Next, l2)
		return l1
	} else {
		l2.Next = merge(l1, l2.Next)
		return l2
	}
}

func sortList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	p, s, f := head, head, head
	for nil != f && nil != f.Next {
		p, s, f = s, s.Next, f.Next.Next
	}
	p.Next = nil
	return merge(sortList(head), sortList(s))
}
