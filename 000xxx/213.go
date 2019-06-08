func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func rob(nums []int) int {
	l := len(nums)
	if 0 == l {
		return 0
	} else if 1 == l {
		return nums[0]
	}
	cal := func(l, r int) int {
		y, n := 0, 0
		for i := l; i < r; i++ {
			y, n = n+nums[i], max(y, n)
		}
		return max(y, n)
	}
	return max(cal(0, l-1), cal(1, l))
}
