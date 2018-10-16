func isPalindrome(head *ListNode) bool {
	ar := []int{}
	for nil != head {
		ar = append(ar, head.Val)
		head = head.Next
	}
	for i := 0; i < (len(ar) >> 1); i++ {
		if ar[i] != ar[len(ar)-1-i] {
			return false
		}
	}
	return true
}
