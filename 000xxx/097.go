func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	t := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		t[i] = make([]bool, n+1)
	}
	t[0][0] = true
	for i := 1; i <= m; i++ {
		t[i][0] = t[i-1][0] && (s1[i-1] == s3[i-1])
	}
	for j := 1; j <= n; j++ {
		t[0][j] = t[0][j-1] && (s2[j-1] == s3[j-1])
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			t[i][j] = (t[i-1][j] && s1[i-1] == s3[i-1+j]) || (t[i][j-1] && s2[j-1] == s3[j-1+i])
		}
	}
	return t[m][n]
}
