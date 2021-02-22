func minOperations(n int) int {
	o := 0
	for i := 0; i < n; i++ {
		if v := 2*i + 1; v > n {
			o += v - n
		}
	}
	return o
}
