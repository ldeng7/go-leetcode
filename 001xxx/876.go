func countGoodSubstrings(s string) int {
	o := 0
	for i, j := 0, 0; i < len(s); i++ {
		if i-j != 2 {
			continue
		}
		if s[j] != s[j+1] && s[i] != s[i-1] && s[i] != s[j] {
			o++
		}
		j++
	}
	return o
}
