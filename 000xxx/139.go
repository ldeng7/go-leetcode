func wordBreak(s string, wordDict []string) bool {
	m := map[string]bool{}
	for _, w := range wordDict {
		m[w] = true
	}
	l := len(s)
	t := make([]bool, l+1)
	t[0] = true
	for i := 0; i <= l; i++ {
		for j := 0; j < i; j++ {
			if t[j] && m[s[j:i]] {
				t[i] = true
				break
			}
		}
	}
	return t[l]
}
