func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{}
	l := h
	c := 0
	for nil != l1 || nil != l2 {
		d1, d2 := 0, 0
		if nil != l1 {
			d1 = l1.Val
		}
		if nil != l2 {
			d2 = l2.Val
		}
		d := d1 + d2 + c
		if d >= 10 {
			d -= 10
			c = 1
		} else {
			c = 0
		}
		l.Next = &ListNode{Val: d}
		l = l.Next
		if nil != l1 {
			l1 = l1.Next
		}
		if nil != l2 {
			l2 = l2.Next
		}
	}
	if 1 == c {
		l.Next = &ListNode{Val: 1}
	}
	return h.Next
}