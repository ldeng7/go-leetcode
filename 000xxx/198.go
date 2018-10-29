func rob(nums []int) int {
	l := len(nums)
	if l >= 3 {
		nums[2] += nums[0]
	}
	for i := 3; i < l; i++ {
		if nums[i-3] > nums[i-2] {
			nums[i] += nums[i-3]
		} else {
			nums[i] += nums[i-2]
		}
	}
	if l >= 2 {
		if nums[l-1] > nums[l-2] {
			return nums[l-1]
		}
		return nums[l-2]
	} else if l == 1 {
		return nums[0]
	}
	return 0
}
