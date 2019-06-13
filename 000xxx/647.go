func countSubstrings(s string) int {
	o, l := 0, len(s)
	t := make([][]bool, l)
	for i := 0; i < l; i++ {
		t[i] = make([]bool, l)
	}
	for i := l - 1; i >= 0; i-- {
		for j := i; j < l; j++ {
			t[i][j] = (s[i] == s[j]) && (j-i <= 2 || t[i+1][j-1])
			if t[i][j] {
				o++
			}
		}
	}
	return o
}
