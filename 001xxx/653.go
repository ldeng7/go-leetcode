func minimumDeletions(s string) int {
	o, c := 0, 0
	for i := 0; i < len(s); i++ {
		if b := s[i]; c > 0 && b == 'a' {
			o++
			c--
		} else if b == 'b' {
			c++
		}
	}
	return o
}
