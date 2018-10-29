func findPaths(m int, n int, N int, i int, j int) int {
	t := make([][][]int, N+1)
	for i := 0; i <= N; i++ {
		t[i] = make([][]int, m)
		for j := 0; j < m; j++ {
			t[i][j] = make([]int, n)
		}
	}
	for i := 1; i <= N; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < n; k++ {
				a, b, c, d := 1, 1, 1, 1
				if j != 0 {
					a = t[i-1][j-1][k]
				}
				if j != m-1 {
					b = t[i-1][j+1][k]
				}
				if k != 0 {
					c = t[i-1][j][k-1]
				}
				if k != n-1 {
					d = t[i-1][j][k+1]
				}
				t[i][j][k] = (a + b + c + d) % 1000000007
			}
		}
	}
	return t[N][i][j]
}
