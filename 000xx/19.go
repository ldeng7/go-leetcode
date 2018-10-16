func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	for nil != head {
		nodes = append(nodes, head)
		head = head.Next
	}
	if len(nodes) <= 1 {
		return nil
	}
	if 1 == n {
		nodes[len(nodes)-2].Next = nil
		return nodes[0]
	} else if n < len(nodes) {
		nodes[len(nodes)-n-1].Next = nodes[len(nodes)-n+1]
		return nodes[0]
	}
	return nodes[1]
}
