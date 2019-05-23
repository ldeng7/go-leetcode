func reverseBetween(head *ListNode, m int, n int) *ListNode {
	h := &ListNode{0, head}
	prev := h
	for i := 0; i < m-1; i++ {
		prev = prev.Next
	}
	cur := prev.Next
	for i := m; i < n; i++ {
		cur.Next, cur.Next.Next, prev.Next = cur.Next.Next, prev.Next, cur.Next
	}
	return h.Next
}
