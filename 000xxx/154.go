func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func findMin(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	o, l, r := nums[0], 0, len(nums)-1
	for l < r-1 {
		m := l + (r-l)>>1
		if nums[l] < nums[m] {
			o, l = min(o, nums[l]), m+1
		} else if nums[l] > nums[m] {
			o, r = min(o, nums[r]), m
		} else {
			l++
		}
	}
	o = min(o, min(nums[l], nums[r]))
	return o
}
