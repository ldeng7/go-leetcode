func isValid(S string) bool {
	l := len(S)
	if l%3 != 0 {
		return false
	}
	s := []byte{}
	for i := 0; i < l; i++ {
		c := S[i]
		if c == 'c' {
			j := len(s) - 2
			if j < 0 || s[j] != 'a' || s[j+1] != 'b' {
				return false
			}
			s = s[:j]
		} else {
			s = append(s, c)
		}
	}
	return true
}
