func shipWithinDays(weights []int, D int) int {
	l, r := 0, 0
	for _, w := range weights {
		if w > l {
			l = w
		}
		r += w
	}
	for l < r {
		m := l + (r-l)>>1
		t, d := 0, 1
		for _, w := range weights {
			t += w
			if t > m {
				d, t = d+1, w
			}
		}
		if d > D {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
