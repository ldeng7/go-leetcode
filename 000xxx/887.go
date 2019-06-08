func superEggDrop(K int, N int) int {
	t := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		t[i] = make([]int, N+1)
	}
	for j := 1; j <= N; j++ {
		for i := 1; i <= K; i++ {
			t[i][j] = t[i][j-1] + t[i-1][j-1] + 1
			if t[i][j] >= N {
				return j
			}
		}
	}
	return N
}
