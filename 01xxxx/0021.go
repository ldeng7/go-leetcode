func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{}
	l := h
	for nil != l1 && nil != l2 {
		if l1.Val <= l2.Val {
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
		l = l.Next
	}
	if l1 != nil {
		l.Next = l1
	} else {
		l.Next = l2
	}
	return h.Next
}
