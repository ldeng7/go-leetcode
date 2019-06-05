func insertionSortList(head *ListNode) *ListNode {
	h := &ListNode{}
	for nil != head {
		t, n := head.Next, h
		for nil != n.Next && n.Next.Val <= head.Val {
			n = n.Next
		}
		head.Next, n.Next, head = n.Next, head, t
	}
	return h.Next
}
