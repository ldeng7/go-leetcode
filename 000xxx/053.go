func maxSubArray(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	o, c := nums[0], 0
	for _, n := range nums {
		if c > 0 {
			c += n
		} else {
			c = n
		}
		if c > o {
			o = c
		}

	}
	return o
}
