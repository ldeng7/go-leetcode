func middleNode(head *ListNode) *ListNode {
	n, c := head, 0
	for nil != n {
		n = n.Next
		c++
	}
	for n, c = head, c>>1; 0 != c; c-- {
		n = n.Next
	}
	return n
}
