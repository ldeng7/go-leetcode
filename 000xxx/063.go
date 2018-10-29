func uniquePathsWithObstacles(grid [][]int) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	if grid[0][0] == 1 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	t := make([]int, n)
	t[0] = 1
	for j := 1; j < n; j++ {
		if grid[0][j] == 1 {
			t[j] = 0
		} else {
			t[j] = t[j-1]
		}
	}
	for i := 1; i < m; i++ {
		if grid[i][0] == 1 {
			t[0] = 0
		}
		for j := 1; j < n; j++ {
			if grid[i][j] == 1 {
				t[j] = 0
			} else {
				t[j] += t[j-1]
			}
		}
	}
	return t[n-1]
}
