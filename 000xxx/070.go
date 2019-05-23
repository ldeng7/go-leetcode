func climbStairs(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	t := make([]int, n)
	t[0], t[1] = 1, 2
	for i := 2; i < n; i++ {
		t[i] = t[i-1] + t[i-2]
	}
	return t[n-1]
}
