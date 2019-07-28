func isPalindrome(head *ListNode) bool {
	ar := []int{}
	for nil != head {
		ar = append(ar, head.Val)
		head = head.Next
	}
	l := len(ar)
	for i := 0; i < l>>1; i++ {
		if ar[i] != ar[l-1-i] {
			return false
		}
	}
	return true
}
