func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func getMaximumGold(grid [][]int) int {
	o, m, n := 0, len(grid), len(grid[0])
	var cal func(int, int, int)
	cal = func(i, j, s int) {
		pg := &grid[i][j]
		g := *pg
		s += g
		o = max(o, s)
		*pg = 0
		if i > 0 && grid[i-1][j] != 0 {
			cal(i-1, j, s)
		}
		if j > 0 && grid[i][j-1] != 0 {
			cal(i, j-1, s)
		}
		if i < m-1 && grid[i+1][j] != 0 {
			cal(i+1, j, s)
		}
		if j < n-1 && grid[i][j+1] != 0 {
			cal(i, j+1, s)
		}
		*pg = g
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				cal(i, j, 0)
			}
		}
	}
	return o
}
