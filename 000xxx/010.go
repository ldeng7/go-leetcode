func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	t := make([][]bool, m+1)
	t[0] = make([]bool, n+1)
	t[0][0] = true
	for j := 2; j <= n; j++ {
		if p[j-1] == '*' {
			t[0][j] = t[0][j-2]
		}
	}
	for i := 1; i <= m; i++ {
		t[i] = make([]bool, n+1)
		if 0 != n {
			t[i][1] = t[i-1][0] && (s[i-1] == p[0] || p[0] == '.')
		}
		for j := 2; j <= n; j++ {
			if p[j-1] == '*' {
				t[i][j] = t[i][j-2] || ((s[i-1] == p[j-2] || p[j-2] == '.') && t[i-1][j])
			} else {
				t[i][j] = t[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	return t[m][n]
}
