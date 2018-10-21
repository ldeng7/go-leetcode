func do(grid [][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) {
		return 0
	}
	if grid[y][x] == 0 {
		return 0
	}
	out := 1
	grid[y][x] = 0
	out += do(grid, x-1, y)
	out += do(grid, x, y-1)
	out += do(grid, x+1, y)
	out += do(grid, x, y+1)
	return out
}

func maxAreaOfIsland(grid [][]int) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	out := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			a := do(grid, j, i)
			if a > out {
				out = a
			}
		}
	}
	return out
}
