func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h := &ListNode{}
	n, c := h, 0
	for nil != l1 || nil != l2 {
		d1, d2 := 0, 0
		if nil != l1 {
			d1, l1 = l1.Val, l1.Next
		}
		if nil != l2 {
			d1, l2 = l2.Val, l2.Next
		}
		d := d1 + d2 + c
		if d >= 10 {
			d -= 10
			c = 1
		} else {
			c = 0
		}
		n.Next = &ListNode{Val: d}
		n = n.Next
	}
	if 1 == c {
		n.Next = &ListNode{Val: 1}
	}
	return h.Next
}
