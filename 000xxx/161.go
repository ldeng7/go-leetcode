func isOneEditDistance(s string, t string) bool {
	if len(s) == len(t) {
		c := 0
		for i := 0; i < len(s) && c <= 1; i++ {
			if s[i] != t[i] {
				c++
			}
		}
		return c == 1
	} else {
		if len(t) > len(s) {
			s, t = t, s
		}
		if len(s)-len(t) != 1 {
			return false
		}
		c := 0
		for i, j := 0, 0; j < len(t) && c <= 1; i++ {
			if s[i] != t[j] {
				c++
			} else {
				j++
			}
		}
		return c <= 1
	}
}
