func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxAscendingSum(nums []int) int {
	t := nums[0]
	o, p := t, t
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		if n > p {
			t += n
		} else {
			t = n
		}
		o, p = max(t, o), n
	}
	return o
}
