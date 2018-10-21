func uniquePaths(m int, n int) int {
	t := make([]int, n)
	for j := 0; j < n; j++ {
		t[j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			t[j] += t[j-1]
		}
	}
	return t[n-1]
}
