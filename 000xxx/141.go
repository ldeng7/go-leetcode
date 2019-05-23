func hasCycle(head *ListNode) bool {
	if nil == head {
		return false
	}
	s, f := head, head
	for nil != f && nil != f.Next {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}
	return false
}
