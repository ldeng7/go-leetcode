func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minOperations(nums []int) int {
	o, t := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		t = max(n, t+1)
		o += t - n
	}
	return o
}
