func insertionSortList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	h := &ListNode{Next: head}
	t, n, m := h, head, head.Next
	for nil != m {
		if m.Val < n.Val {
			t = h
			for t.Next.Val < m.Val {
				t = t.Next
			}
			t.Next, n.Next, m.Next = m, m.Next, t.Next
			m = n.Next
		} else {
			n, m = m, m.Next
		}
	}
	return h.Next
}
