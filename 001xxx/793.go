func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maximumScore(nums []int, k int) int {
	o, n := 0, len(nums)
	for i := k - 1; i >= 0; i-- {
		nums[i] = min(nums[i], nums[i+1])
	}
	for i := k + 1; i < n; i++ {
		nums[i] = min(nums[i], nums[i-1])
	}
	l, r := k, k
	for i := nums[k]; i >= 0; i-- {
		for ; l-1 >= 0 && nums[l-1] >= i; l-- {
		}
		for ; r+1 < n && nums[r+1] >= i; r++ {
		}
		o = max(o, (r-l+1)*i)
	}
	return o
}
