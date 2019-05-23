func countSegments(s string) int {
	i, c := 0, 0
	for ; i < len(s); i++ {
		if ' ' == s[i] {
			continue
		}
		c++
		for ; i < len(s) && ' ' != s[i]; i++ {
		}
	}
	return c
}
