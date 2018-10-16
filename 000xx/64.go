func minPathSum(grid [][]int) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	m, n := len(grid), len(grid[0])
	t := make([][]int, m)
	for i := 0; i < m; i++ {
		t[i] = make([]int, n)
	}

	t[0][0] = grid[0][0]
	for j := 1; j < n; j++ {
		t[0][j] = grid[0][j] + t[0][j-1]
	}
	for i := 1; i < m; i++ {
		t[i][0] = grid[i][0] + t[i-1][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			t[i][j] = grid[i][j]
			if t[i-1][j] < t[i][j-1] {
				t[i][j] += t[i-1][j]
			} else {
				t[i][j] += t[i][j-1]
			}
		}
	}
	return t[m-1][n-1]
}
