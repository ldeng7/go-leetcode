func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	t := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		t[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		t[i][0] = i
	}
	for j := 0; j <= n; j++ {
		t[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				t[i][j] = t[i-1][j-1]
			} else {
				t[i][j] = min(t[i-1][j-1], min(t[i-1][j], t[i][j-1])) + 1
			}
		}
	}
	return t[m][n]
}
