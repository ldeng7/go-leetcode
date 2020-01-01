func getDecimalValue(head *ListNode) int {
	o := 0
	for n := head; nil != n; n = n.Next {
		o = (o << 1) | n.Val
	}
	return o
}
