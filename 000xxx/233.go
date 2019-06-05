func countDigitOne(n int) int {
	o, a, b := 0, 1, 1
	for n > 0 {
		o += (n + 8) / 10 * a
		if n%10 == 1 {
			o += b
		}
		b += n % 10 * a
		a *= 10
		n /= 10
	}
	return o
}
