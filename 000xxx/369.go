func cal(n *ListNode) int {
	if nil == n {
		return 1
	}
	sum := n.Val + cal(n.Next)
	n.Val = sum % 10
	return sum / 10
}

func plusOne(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	if 1 == cal(head) {
		return &ListNode{Val: 1, Next: head}
	}
	return head
}
