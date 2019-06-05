func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func wiggleMaxLength(nums []int) int {
	a, b, l := 1, 1, len(nums)
	for i := 1; i < l; i++ {
		if nums[i] > nums[i-1] {
			a = b + 1
		} else if nums[i] < nums[i-1] {
			b = a + 1
		}
	}
	return min(l, max(a, b))
}
