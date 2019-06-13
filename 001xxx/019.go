func nextLargerNodes(head *ListNode) []int {
	out, st, i := []int{}, []int{}, 0
	for nil != head {
		for 0 != len(st) {
			t := st[len(st)-1]
			if head.Val <= out[t] {
				break
			}
			out[t], st = head.Val, st[:len(st)-1]
		}
		out, st = append(out, head.Val), append(st, i)
		head, i = head.Next, i+1
	}

	for i := len(st) - 1; i >= 0; i-- {
		out[st[i]] = 0
	}
	return out
}
