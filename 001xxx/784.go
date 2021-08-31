func checkOnesSegment(s string) bool {
	l := len(s)
	p1, p2 := s[0], s[l-1]
	for i, j := 1, l-2; i <= j; i, j = i+1, j-1 {
		c1, c2 := s[i], s[j]
		if (p1 == '0' && c1 == '1') || (c2 == '0' && p2 == '1') {
			return false
		}
		p1, p2 = c1, c2
	}
	return true
}
