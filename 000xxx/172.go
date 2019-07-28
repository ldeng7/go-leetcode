func trailingZeroes(n int) int {
	o := 0
	for n > 0 {
		o += n / 5
		n /= 5
	}
	return o
}
