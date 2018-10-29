func divide(dividend int, divisor int) int {
	out := dividend / divisor
	if out < -2147483648 {
		out = -2147483648
	} else if out > 2147483647 {
		out = 2147483647
	}
	return out
}