func findPeakElement(nums []int) int {
	l, h := 0, len(nums)-1
	for l < h {
		m := int((uint(l) + uint(h)) >> 1)
		if nums[m] < nums[m+1] {
			l = m + 1
		} else {
			h = m
		}
	}
	return l
}
