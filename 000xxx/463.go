func islandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	c := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i == 0 || grid[i-1][j] == 0 {
				c++
			}
			if i == m-1 || grid[i+1][j] == 0 {
				c++
			}
			if j == 0 || grid[i][j-1] == 0 {
				c++
			}
			if j == n-1 || grid[i][j+1] == 0 {
				c++
			}
		}
	}
	return c
}
