func kLengthApart(nums []int, k int) bool {
	p := -100001
	for i, n := range nums {
		if n == 1 {
			if i-p < k+1 {
				return false
			}
			p = i
		}
	}
	return true
}
