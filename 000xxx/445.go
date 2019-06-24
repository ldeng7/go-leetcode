func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := []int{}, []int{}
	for nil != l1 {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for nil != l2 {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	n, o := 0, &ListNode{Val: 0}
	for {
		ls1, ls2 := len(s1), len(s2)
		if 0 == ls1 && 0 == ls2 {
			break
		}
		if 0 != ls1 {
			n, s1 = n+s1[ls1-1], s1[:ls1-1]
		}
		if 0 != ls2 {
			n, s2 = n+s2[ls2-1], s2[:ls2-1]
		}
		pn := &ListNode{Next: o}
		if n < 10 {
			o.Val, n = n, 0
		} else {
			o.Val, pn.Val, n = n-10, 1, 1
		}
		o = pn
	}
	if o.Val != 0 {
		return o
	}
	return o.Next
}
