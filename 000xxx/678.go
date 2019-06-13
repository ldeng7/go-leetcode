func checkValidString(s string) bool {
	l, h := 0, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' {
			l, h = l+1, h+1
		} else {
			if l > 0 {
				l--
			}
			if c == ')' {
				h--
			} else {
				h++
			}
		}
		if h < 0 {
			return false
		}
	}
	return l == 0
}
