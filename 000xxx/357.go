func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	o, c := 10, 9
	for i := 2; i <= n; i++ {
		c *= (11 - i)
		o += c
	}
	return o
}
