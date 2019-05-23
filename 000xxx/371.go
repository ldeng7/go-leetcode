func getSum(a int, b int) int {
	if 0 == b {
		return a
	}
	return getSum(a^b, (a&b)<<1)
}
