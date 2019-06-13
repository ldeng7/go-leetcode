func checkRecord(n int) int {
	p, pl := make([]int, n+1), make([]int, n+1)
	p[0], p[1], pl[0], pl[1] = 1, 1, 1, 2
	for i := 2; i <= n; i++ {
		p[i] = pl[i-1]
		pl[i] = (p[i] + p[i-1] + p[i-2]) % 1000000007
	}
	o := pl[n]
	for i := 0; i < n; i++ {
		o += pl[i] * pl[n-1-i]
		o %= 1000000007
	}
	return o
}
