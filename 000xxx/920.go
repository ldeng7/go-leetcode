func numMusicPlaylists(N int, L int, K int) int {
	b := 1
	for i := 1; i <= N; i++ {
		b = b * i % 1000000007
	}
	t := make([][]int, L-N+1)
	for i := 0; i <= L-N; i++ {
		t[i] = make([]int, N-K)
		for j := 0; j < N-K; j++ {
			t[i][j] = 1
		}
	}

	for i := 1; i <= L-N; i++ {
		for j := 1; j < N-K; j++ {
			t[i][j] = (t[i][j-1] + (j+1)*t[i-1][j]) % 1000000007
		}
	}
	return t[L-N][N-K-1] * b % 1000000007
}
