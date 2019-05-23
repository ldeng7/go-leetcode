func minPathSum(grid [][]int) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	m, n := len(grid), len(grid[0])
	t := make([]int, n)
	copy(t, grid[0])
	for j := 1; j < n; j++ {
		t[j] += t[j-1]
	}

	for i := 1; i < m; i++ {
		t[0] += grid[i][0]
		for j := 1; j < n; j++ {
			if t[j-1] < t[j] {
				t[j] = t[j-1] + grid[i][j]
			} else {
				t[j] += grid[i][j]
			}
		}
	}
	return t[n-1]
}
