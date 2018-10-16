func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * ((n >> 1) + 1)
	if n&1 == 0 {
		sum -= n >> 1
	}
	s := 0
	for _, num := range nums {
		s += num
	}
	return sum - s
}
