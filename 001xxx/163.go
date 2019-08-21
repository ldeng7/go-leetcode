func lastSubstring(s string) string {
	l, m := len(s), 0
	for i := 1; i < l; i++ {
		j := 0
		for ; j < l-i; j++ {
			a, b := s[m+j], s[i+j]
			if a == b {
				continue
			}
			if a < b {
				m = i
			}
			break
		}
		if i+j == l {
			break
		}
	}
	return s[m:]
}
