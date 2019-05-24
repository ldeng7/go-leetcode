func partition(head *ListNode, x int) *ListNode {
	hl, hg := &ListNode{0, nil}, &ListNode{0, nil}
	n, l, g := head, hl, hg
	for nil != n {
		if n.Val >= x {
			g.Next = n
			g = n
		} else {
			l.Next = n
			l = n
		}
		n = n.Next
	}
	g.Next = nil
	if g != hg {
		l.Next = hg.Next
	}
	return hl.Next
}
