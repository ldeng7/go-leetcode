func plusOne(digits []int) []int {
	acc := 1
	for i := len(digits) - 1; i >= 0 && acc == 1; i-- {
		n := digits[i] + 1
		if n == 10 {
			n = 0
		} else {
			acc = 0
		}
		digits[i] = n
	}
	if acc == 0 {
		return digits
	}
	return append([]int{1}, digits...)
}
