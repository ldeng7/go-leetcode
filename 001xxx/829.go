func getMaximumXor(nums []int, maximumBit int) []int {
	l, m := len(nums), (1<<maximumBit)-1
	for i := 0; i < l; i++ {
		m ^= nums[i]
		nums[i] = m
	}
	for i := 0; i < l>>1; i++ {
		nums[i], nums[l-i-1] = nums[l-i-1], nums[i]
	}
	return nums
}
