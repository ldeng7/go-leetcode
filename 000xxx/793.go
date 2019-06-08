func preimageSizeFZF(K int) int {
	l, r := 0, 5*(K+1)
	for l < r {
		m := l + (r-l)>>1
		c := 0
		for i := 5; m/i > 0; i *= 5 {
			c += m / i
		}
		if c == K {
			return 5
		} else if c < K {
			l = m + 1
		} else {
			r = m
		}
	}
	return 0
}
