func rotateRight(head *ListNode, k int) *ListNode {
	if nil == head {
		return nil
	}
	n, c := head, 1
	for ; n.Next != nil; n = n.Next {
		c++
	}
	n.Next = head
	n = head
	for i := 0; i < c-k%c-1; i++ {
		n = n.Next
	}
	h := n.Next
	n.Next = nil
	return h
}
