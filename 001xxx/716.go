func totalMoney(n int) int {
	o := 0
	for i := 0; i < n; i++ {
		o += i/7 + i%7 + 1
	}
	return o
}
