func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		if nums[m] == nums[m^1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
