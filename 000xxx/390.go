func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}
	return (n>>1 - lastRemaining(n>>1) + 1) << 1
}
