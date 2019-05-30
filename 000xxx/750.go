func countCornerRectangles(grid [][]int) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	m, n := len(grid), len(grid[0])
	out := 0
	for i := 0; i < m-1; i++ {
		uno := []int{}
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				uno = append(uno, j)
			}
		}
		for j := i + 1; j < m; j++ {
			c := 0
			for _, k := range uno {
				if grid[j][k] != 0 {
					c++
				}
			}
			out += c * (c - 1) / 2
		}
	}
	return out
}
