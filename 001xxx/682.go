func longestPalindromeSubseq(s string) int {
	l := len(s)
	t := make([][][2]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([][2]int, l)
	}
	for i := l - 2; i >= 0; i-- {
		for j := i + 1; j < l; j++ {
			if p := t[i+1][j-1]; s[i] == s[j] && s[i] != byte(p[1]) {
				t[i][j] = [2]int{p[0] + 2, int(s[i])}
			} else if p = t[i+1][j]; p[0] >= t[i][j-1][0] {
				t[i][j] = p
			} else {
				t[i][j] = t[i][j-1]
			}
		}
	}
	return t[0][l-1][0]
}
