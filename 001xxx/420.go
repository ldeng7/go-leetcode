func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func numOfArrays(n int, m int, k int) int {
	if k == 0 || k > m {
		return 0
	}
	t, s := make([][][]int, n+1), make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		art, ars := make([][]int, k+1), make([][]int, k+1)
		for j := 0; j <= k; j++ {
			art[j], ars[j] = make([]int, m+1), make([]int, m+1)
		}
		t[i], s[i] = art, ars
	}
	for i := 1; i <= m; i++ {
		t[1][1][i], s[1][1][i] = 1, i
	}

	for i := 2; i <= n; i++ {
		for j, je := 1, min(i, k); j <= je; j++ {
			for p := j; p <= m; p++ {
				t[i][j][p] = (t[i-1][j][p]*p + s[i-1][j-1][p-1]) % 1000000007
				s[i][j][p] = (s[i][j][p-1] + t[i][j][p]) % 1000000007
			}
		}
	}
	return s[n][k][m]
}
