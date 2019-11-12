func maximizeSweetness(sweetness []int, K int) int {
	o, s := 0, 0
	for _, sw := range sweetness {
		s += sw
	}
	if K == 0 {
		return s
	}
	l, r := 0, s/K+1
	for l <= r {
		m := l + (r-l)>>1
		s, n := 0, 0
		for _, sw := range sweetness {
			s += sw
			if s >= m {
				s, n = 0, n+1
			}
		}
		if n >= K+1 {
			o, l = m, m+1
		} else {
			r = m - 1
		}
	}
	return o
}
