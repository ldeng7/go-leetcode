func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxProduct(nums []int) int {
	o := nums[0]
	mx, mn := o, o
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			mx, mn = mn, mx
		}
		mx = max(nums[i], mx*nums[i])
		mn = min(nums[i], mn*nums[i])
		o = max(o, mx)
	}
	return o
}
