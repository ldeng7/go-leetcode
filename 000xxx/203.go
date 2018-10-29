func removeElements(head *ListNode, val int) *ListNode {
	root := &ListNode{}
	prev := root
	for nil != head {
		if head.Val == val {
			head = head.Next
			continue
		}
		prev.Next = head
		prev = head
		head = head.Next
		prev.Next = nil
	}
	return root.Next
}
