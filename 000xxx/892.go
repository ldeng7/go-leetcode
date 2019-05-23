func surfaceArea(grid [][]int) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	sum := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		sum += grid[i][0]
		sum += grid[i][n-1]
	}
	for j := 0; j < n; j++ {
		sum += grid[0][j]
		sum += grid[m-1][j]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				sum += 2
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			d := grid[i][j] - grid[i-1][j]
			if d < 0 {
				d = -d
			}
			sum += d
		}
	}
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			d := grid[i][j] - grid[i][j-1]
			if d < 0 {
				d = -d
			}
			sum += d
		}
	}
	return sum
}
