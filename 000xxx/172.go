func trailingZeroes(n int) int {
	c := 0
	for n > 0 {
		c += n / 5
		n /= 5
	}
	return c
}
