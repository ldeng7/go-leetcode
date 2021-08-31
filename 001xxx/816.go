func truncateSentence(s string, k int) string {
	i, l := 0, len(s)
	for n := 0; i < l; i++ {
		if s[i] == ' ' {
			n++
			if n >= k {
				break
			}
		}
	}
	return s[:i]
}
