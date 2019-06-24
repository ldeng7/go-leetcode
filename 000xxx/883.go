func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func projectionArea(grid [][]int) int {
	o, l := 0, len(grid)
	for i := 0; i < l; i++ {
		r, c := 0, 0
		for j := 0; j < l; j++ {
			if grid[i][j] > 0 {
				o++
			}
			r = max(r, grid[i][j])
			c = max(c, grid[j][i])
		}
		o += r + c
	}
	return o
}
