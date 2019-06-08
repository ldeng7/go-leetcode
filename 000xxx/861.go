func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])
	o := (1 << uint(n-1)) * m
	for j := 1; j < n; j++ {
		c := 0
		for i := 0; i < m; i++ {
			if A[i][j] == A[i][0] {
				c++
			}
		}
		o += max(c, m-c) * (1 << uint(n-1-j))
	}
	return o
}
