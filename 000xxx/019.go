func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	for nil != head {
		nodes = append(nodes, head)
		head = head.Next
	}
	l := len(nodes)
	if l <= 1 {
		return nil
	}
	if 1 == n {
		nodes[l-2].Next = nil
		return nodes[0]
	} else if n < l {
		nodes[l-n-1].Next = nodes[l-n+1]
		return nodes[0]
	}
	return nodes[1]
}
