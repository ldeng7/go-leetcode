func numWays(n int, k int) int {
	if n == 0 {
		return 0
	}
	s, d := 0, k
	for i := 2; i <= n; i++ {
		s, d = d, (s+d)*(k-1)
	}
	return s + d
}
