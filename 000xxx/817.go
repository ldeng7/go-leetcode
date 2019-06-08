func numComponents(head *ListNode, G []int) int {
	o := 0
	s := map[int]bool{}
	for _, g := range G {
		s[g] = true
	}
	for nil != head {
		if s[head.Val] && (nil == head.Next || !s[head.Next.Val]) {
			o++
		}
		head = head.Next
	}
	return o
}
