func projectionArea(grid [][]int) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	sum := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		h := 0
		for j := 0; j < n; j++ {
			x := grid[i][j]
			if x != 0 {
				sum++ // z-axis
			}
			if x > h {
				h = x
			}
		}
		sum += h // x-axis
	}
	for j := 0; j < n; j++ {
		h := 0
		for i := 0; i < m; i++ {
			x := grid[i][j]
			if x > h {
				h = x
			}
		}
		sum += h // y-axis
	}
	return sum
}
