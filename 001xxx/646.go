func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}
	ar := make([]int, n+1)
	ar[0], ar[1] = 0, 1
	o := 1
	for i := 2; i <= n; i++ {
		if i&1 == 0 {
			ar[i] = ar[i>>1]
		} else {
			ar[i] = ar[i>>1] + ar[i>>1+1]
		}
		o = max(o, ar[i])
	}
	return o
}
