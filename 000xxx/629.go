func kInversePairs(n int, k int) int {
	t := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		t[i] = make([]int, k+1)
	}
	t[0][0] = 1
	for i := 1; i <= n; i++ {
		t[i][0] = 1
		for j := 1; j <= k; j++ {
			t[i][j] = t[i-1][j] + t[i][j-1]
			if j >= i {
				t[i][j] -= t[i-1][j-i] - 1000000007
			}
			t[i][j] %= 1000000007
		}
	}
	return t[n][k]
}
