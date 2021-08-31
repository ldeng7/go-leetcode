func secondHighest(s string) int {
	a, b, l := -1, -1, len(s)
	bs := [10]bool{}
	for i := 0; i < l; i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			continue
		}
		c := int(ch - '0')
		if bs[c] {
			continue
		}
		bs[c] = true
		if c > b {
			b = c
			if b > a {
				a, b = b, a
			}
		}
	}
	return b
}
