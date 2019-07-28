func rangeBitwiseAnd(m int, n int) int {
	if 0 == m {
		return 0
	}
	t := 1
	for m != n {
		m, n, t = m>>1, n>>1, t<<1
	}
	return m * t
}
