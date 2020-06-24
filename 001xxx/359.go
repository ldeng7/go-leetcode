func countOrders(n int) int {
	n <<= 1
	o := 1
	for i := 1; i <= n; i++ {
		o *= i
		if i&1 == 0 {
			o >>= 1
		}
		o %= 1000000007
	}
	return o
}
