func reorderList(head *ListNode) {
	if nil == head || nil == head.Next || nil == head.Next.Next {
		return
	}
	n, st := head, []*ListNode{}
	for nil != n {
		st = append(st, n)
		n = n.Next
	}
	c := (len(st) - 1) >> 1
	n = head
	for ; c > 0; c-- {
		t := st[len(st)-1]
		st = st[:len(st)-1]
		next := n.Next
		n, t.Next, n.Next = next, next, t
	}
	st[len(st)-1].Next = nil
}
