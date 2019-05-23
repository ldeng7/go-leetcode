func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	t := make([][]bool, ls+1)
	for i := 0; i <= ls; i++ {
		t[i] = make([]bool, lp+1)
	}
	t[0][0] = true
	for i := 1; i <= lp; i++ {
		if '*' == p[i-1] {
			t[0][i] = t[0][i-1]
		}
	}
	for i := 1; i <= ls; i++ {
		for j := 1; j <= lp; j++ {
			if '*' == p[j-1] {
				t[i][j] = t[i-1][j] || t[i][j-1]
			} else {
				t[i][j] = (s[i-1] == p[j-1] || '?' == p[j-1]) && t[i-1][j-1]
			}
		}
	}
	return t[ls][lp]
}
