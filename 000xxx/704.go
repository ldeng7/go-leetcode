func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		if nums[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	if nums[l] != target {
		return -1
	}
	return l
}
