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

func mergeKLists(lists []*ListNode) *ListNode {
	if 0 == len(lists) {
		return nil
	} else if 1 == len(lists) {
		return lists[0]
	}
	h := lists[0]
	for _, l := range lists[1:] {
		h = mergeTwoLists(h, l)
	}
	return h
}
