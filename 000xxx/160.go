func getIntersectionNode(headA, headB *ListNode) *ListNode {
	na, nb := headA, headB
	la, lb := 0, 0
	for nil != na {
		na = na.Next
		la++
	}
	for nil != nb {
		nb = nb.Next
		lb++
	}
	na, nb = headA, headB
	if la >= lb {
		d := la - lb
		for i := 0; i < d; i++ {
			na = na.Next
		}
	} else {
		d := lb - la
		for i := 0; i < d; i++ {
			nb = nb.Next
		}
	}
	for nil != na {
		if na == nb {
			return na
		}
		na, nb = na.Next, nb.Next
	}
	return nil
}
