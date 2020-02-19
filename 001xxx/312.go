func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minInsertions(s string) int {
	l := len(s)
	t := make([]int, l)
	for i := l - 2; i >= 0; i-- {
		for j, k := i+1, t[i]; j < l; j++ {
			p := t[j]
			if s[i] == s[j] {
				t[j] = k
			} else {
				t[j] = min(t[j], t[j-1]) + 1
			}
			k = p
		}
	}
	return t[l-1]
}
