func rangeBitwiseAnd(m int, n int) int {
	for m < n {
		n &= (n - 1)
	}
	return n
}
