func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	c := 0
	l, p := 0, 1
	for i := 0; i < len(nums); i++ {
		p *= nums[i]
		for p >= k {
			p /= nums[l]
			l++
		}
		c += i - l + 1
	}
	return c
}
