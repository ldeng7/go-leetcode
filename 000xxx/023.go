func merge2Lists(l1, l2 *ListNode) *ListNode {
	n := &ListNode{}
	h := n
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l1, n.Next = l1.Next, l1
		} else {
			n.Next = l2
			l1, l2 = n.Next.Next, l1
		}
		n = n.Next
	}
	if l1 != nil {
		n.Next = l1
	}
	if l2 != nil {
		n.Next = l2
	}
	return h.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0 {
		return nil
	} else if l == 1 {
		return lists[0]
	}
	m := l >> 1
	return merge2Lists(mergeKLists(lists[:m]), mergeKLists(lists[m:]))
}
