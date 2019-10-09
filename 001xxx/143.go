func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	t := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		t[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				t[i+1][j+1] = t[i][j] + 1
			} else {
				t[i+1][j+1] = max(t[i+1][j], t[i][j+1])
			}
		}
	}
	return t[m][n]
}
