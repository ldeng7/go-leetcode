func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = 1
	}
	for i := 1; i <= n; i++ {
		dp[i][0] = 0
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = dp[i][j-1]
			if t[i-1] == s[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}
	return dp[n][m]
}
