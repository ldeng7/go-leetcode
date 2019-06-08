func numTilings(N int) int {
	if 1 == N {
		return 1
	}
	t := make([]int, N+1)
	t[0], t[1], t[2] = 1, 1, 2
	for i := 3; i <= N; i++ {
		t[i] = (t[i-1]*2 + t[i-3]) % 1000000007
	}
	return t[N]
}
