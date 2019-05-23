func detectCycle(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	s, f := head, head
	for nil != f && nil != f.Next {
		s = s.Next
		f = f.Next.Next
		if s == f {
			break
		}
	}
	if nil == f || nil == f.Next {
		return nil
	}
	s = head
	for s != f {
		s, f = s.Next, f.Next
	}
	return s
}
