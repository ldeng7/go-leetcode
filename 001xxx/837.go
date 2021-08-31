func sumBase(n int, k int) int {
	o := 0
	for ; n > 0; n /= k {
		o += n % k
	}
	return o
}
