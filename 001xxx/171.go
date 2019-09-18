func removeZeroSumSublists(head *ListNode) *ListNode {
	r := &ListNode{0, head}
	m := map[int]*ListNode{0: r}
	s := 0
	for n := head; nil != n; n = n.Next {
		s += n.Val
		if n1, ok := m[s]; ok {
			for p, k := n1, s; p != n; {
				p = p.Next
				k += p.Val
				if p != n {
					delete(m, k)
				}
			}
			n1.Next = n.Next
		} else {
			m[s] = n
		}
	}
	return r.Next
}
