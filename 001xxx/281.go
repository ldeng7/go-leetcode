func subtractProductAndSum(n int) int {
	m, s := 1, 0
	for ; n != 0; n /= 10 {
		d := n % 10
		m *= d
		s += d
	}
	return m - s
}
