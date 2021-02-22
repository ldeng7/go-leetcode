func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxProductPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	tmin, tmax := make([][]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		tmin[i], tmax[i] = make([]int, n), make([]int, n)
	}
	tmin[0][0], tmax[0][0] = grid[0][0], grid[0][0]
	for i := 1; i < m; i++ {
		t := tmin[i-1][0] * grid[i][0]
		tmin[i][0], tmax[i][0] = t, t
	}
	for j := 1; j < n; j++ {
		t := tmin[0][j-1] * grid[0][j]
		tmin[0][j], tmax[0][j] = t, t
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if v := grid[i][j]; v < 0 {
				tmin[i][j] = max(tmax[i-1][j], tmax[i][j-1]) * v
				tmax[i][j] = min(tmin[i-1][j], tmin[i][j-1]) * v
			} else {
				tmin[i][j] = min(tmin[i-1][j], tmin[i][j-1]) * v
				tmax[i][j] = max(tmax[i-1][j], tmax[i][j-1]) * v
			}
		}
	}
	o := tmax[m-1][n-1] % 1000000007
	if o < 0 {
		return -1
	}
	return o
}
