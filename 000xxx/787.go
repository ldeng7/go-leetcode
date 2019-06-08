func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	t := make([][]int, K+2)
	for i := 0; i < K+2; i++ {
		t[i] = make([]int, n)
		for j := 0; j < n; j++ {
			t[i][j] = 1e9
		}
	}
	t[0][src] = 0
	for i := 1; i < K+2; i++ {
		t[i][src] = 0
		for _, f := range flights {
			t1 := t[i-1][f[0]] + f[2]
			if t1 < t[i][f[1]] {
				t[i][f[1]] = t1
			}
		}
	}
	if t[K+1][dst] == 1e9 {
		return -1
	}
	return t[K+1][dst]
}
