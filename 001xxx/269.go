func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func numWays(steps int, arrLen int) int {
	t := [502]int{1}
	for i := 1; i <= steps; i++ {
		m := min(min(arrLen, i+1), steps-i+1)
		for j, k := 0, 0; j < m; j++ {
			t[j], k = k, t[j]
			t[j] = (t[j] + t[j+1] + k) % 1000000007
		}
	}
	return t[0]
}
