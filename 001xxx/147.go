func longestDecomposition(text string) int {
	l := len(text)
	if 0 == l {
		return 0
	}
	for i := 1; i <= l>>1; i++ {
		if text[:i] == text[l-i:] {
			return 2 + longestDecomposition(text[i:l-i])
		}
	}
	return 1
}
