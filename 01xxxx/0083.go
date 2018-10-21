func deleteDuplicates(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	i, j := head.Next, head
	for ; i != nil; i = i.Next {
		if i.Val == j.Val {
			continue
		}
		j = j.Next
		j.Val = i.Val
	}
	j.Next = nil
	return head
}
