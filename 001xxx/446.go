func maxPower(s string) int {
	o, l, p, c := 0, len(s), s[0], 1
	for i := 1; i < l; i++ {
		if b := s[i]; b != p {
			if c > o {
				o = c
			}
			p, c = b, 1
		} else {
			c++
		}
	}
	if c > o {
		o = c
	}
	return o
}
