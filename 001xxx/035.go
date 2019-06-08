func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxUncrossedLines(A []int, B []int) int {
	m, n := len(A), len(B)
	t := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		t[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if A[i-1] == B[j-1] {
				t[i][j] = t[i-1][j-1] + 1
			} else {
				t[i][j] = max(t[i][j-1], t[i-1][j])
			}
		}
	}
	return t[m][n]
}
