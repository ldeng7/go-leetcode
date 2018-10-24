func pivotIndex(nums []int) int {
	if 0 == len(nums) {
		return -1
	}
	l, r := 0, 0
	for _, n := range nums[1:] {
		r += n
	}
	if l == r {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		l += nums[i-1]
		r -= nums[i]
		if l == r {
			return i
		}
	}
	return -1
}
