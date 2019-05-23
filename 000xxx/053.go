func maxSubArray(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	s := nums[0]
	c := 0
	for _, n := range nums {
		if c > 0 {
			c += n
		} else {
			c = n
		}
		if c > s {
			s = c
		}

	}
	return s
}
