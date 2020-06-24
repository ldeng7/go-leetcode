func maxScore(s string) int {
	o, l := 0, len(s)
	var l0, l1, r0, r1 int
	for i := 0; i < l; i++ {
		if s[i] == '0' {
			r0++
		} else {
			r1++
		}
	}
	l--
	for i := 0; i < l; i++ {
		if s[i] == '0' {
			l0++
			r0--
		} else {
			l1++
			r1--
		}
		if t := l0 + r1; t > o {
			o = t
		}
	}
	return o
}
