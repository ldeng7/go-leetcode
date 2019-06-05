func wordBreak(s string, wordDict []string) bool {
	m := map[string]bool{}
	for _, w := range wordDict {
		m[w] = true
	}
	t := make([]bool, len(s)+1)
	t[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if t[j] && m[s[j:i]] {
				t[i] = true
				break
			}
		}
	}
	return t[len(s)]
}
