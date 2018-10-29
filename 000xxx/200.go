func cal(grid [][]byte, x, y int) bool {
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return false
	}
	if grid[y][x] == '0' {
		return false
	}
	grid[y][x] = '0'
	cal(grid, x-1, y)
	cal(grid, x, y-1)
	cal(grid, x+1, y)
	cal(grid, x, y+1)
	return true
}

func numIslands(grid [][]byte) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	out := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if cal(grid, j, i) {
				out++
			}
		}
	}
	return out
}
