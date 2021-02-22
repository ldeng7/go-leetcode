func minInsertions(s string) int {
	o, c, l := 0, 0, len(s)
	for i := 0; i < l; {
		if b := s[i]; b == '(' {
			c++
			i++
		} else {
			if c > 0 {
				c--
			} else {
				o++
			}
			if i < l-1 && s[i+1] == ')' {
				i += 2
			} else {
				o++
				i++
			}
		}
	}
	o += c * 2
	return o
}
