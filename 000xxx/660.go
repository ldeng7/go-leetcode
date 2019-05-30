func newInteger(n int) int {
	o, b := 0, 1
	for n > 0 {
		o += n % 9 * b
		n /= 9
		b *= 10
	}
	return o
}
