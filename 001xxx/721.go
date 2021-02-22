func swapNodes(head *ListNode, k int) *ListNode {
	c, k := 0, k-1
	for n := head; nil != n; n = n.Next {
		c++
	}
	i, j := 0, c-k-1
	if k == j {
		return head
	}
	if j < k {
		j, k = k, j
	}
	n := head
	for ; i < k; i++ {
		n = n.Next
	}
	n1 := n
	for ; i < j; i++ {
		n1 = n1.Next
	}
	n.Val, n1.Val = n1.Val, n.Val
	return head
}
