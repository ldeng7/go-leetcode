func numTrees(n int) int {
	t := make([]int, n+1)
	t[0] = 1
	t[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			t[i] += t[j] * t[i-j-1]
		}
	}
	return t[n]
}
