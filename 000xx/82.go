func deleteDuplicates(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	h := &ListNode{Next: head}
	i, j := head, h
	var v int
	for ; i != nil; i = i.Next {
		if (i != head && i.Val == v) || (i.Next != nil && i.Next.Val == i.Val) {
			v = i.Val
			continue
		}
		j = j.Next
		j.Val = i.Val
	}
	j.Next = nil
	return h.Next
}
