func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	n1, n2 := list1, list2
	for c := 0; c < b+1; c++ {
		if c == a-1 {
			t := n1.Next
			n1.Next = n2
			n1 = t
		} else {
			n1 = n1.Next
		}
	}
	for nil != n2.Next {
		n2 = n2.Next
	}
	n2.Next = n1
	return list1
}
