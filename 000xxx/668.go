func findKthNumber(m int, n int, k int) int {
	l, r := 1, m*n
	for l < r {
		h, c, i, j := l+(r-l)>>1, 0, m, 1
		for i >= 1 && j <= n {
			t := j
			if h > n*i {
				j = n + 1
			} else {
				j = h/i + 1
			}
			c, i = c+(j-t)*i, h/j
		}
		if c < k {
			l = h + 1
		} else {
			r = h
		}
	}
	return r
}
