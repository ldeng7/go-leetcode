func isSubsequence(s string, t string) bool {
	i, ls, lt := 0, len(s), len(t)
	for j := 0; j < lt && i < ls; j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == ls
}
