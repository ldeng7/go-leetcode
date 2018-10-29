func splitListToParts(root *ListNode, k int) []*ListNode {
	out := make([]*ListNode, k)
	l := 0
	for n := root; nil != n; n = n.Next {
		l++
	}
	avg, ext := l/k, l%k
	for i := 0; i < k && nil != root; i++ {
		out[i] = root
		je := avg
		if i < ext {
			je++
		}
		for j := 1; j < je; j++ {
			root = root.Next
		}
		n := root.Next
		root.Next = nil
		root = n
	}
	return out
}
