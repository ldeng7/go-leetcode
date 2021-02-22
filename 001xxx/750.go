func minimumLength(s string) int {
	i, j := 0, len(s)-1
	for i < j {
		c := s[i]
		if s[i] == s[j] {
			for ; i <= j && s[i] == c; i++ {
			}
			for ; i <= j && s[j] == c; j-- {
			}
		} else {
			break
		}
	}
	if i > j {
		return 0
	}
	return j - i + 1
}
