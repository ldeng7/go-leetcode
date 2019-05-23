func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}
	for {
		if 0 != num&3 {
			return 1 == num
		}
		num >>= 2
	}
}
