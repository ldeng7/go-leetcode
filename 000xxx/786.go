func kthSmallestPrimeFraction(A []int, K int) []int {
	var l, r float64 = 0, 1
	p, q, c, n := 0, 1, 0, len(A)
	for {
		m := l + (r-l)*0.5
		c, p = 0, 0
		for i, j := 0, 0; i < n; i++ {
			for j < n && float64(A[i]) > m*float64(A[j]) {
				j++
			}
			c += n - j
			if j < n && p*A[j] < q*A[i] {
				p, q = A[i], A[j]
			}
		}
		if c == K {
			return []int{p, q}
		} else if c < K {
			l = m
		} else {
			r = m
		}
	}
}
