func findDerangement(n int) int {
	out := 1
	for i := 1; i <= n; i++ {
		out = (i*out - ((i & 1) << 1) + 1) % 1000000007
	}
	return out
}
