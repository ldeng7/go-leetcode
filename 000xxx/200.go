func numIslands(grid [][]byte) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	m, n := len(grid), len(grid[0])
	o := 0
	var cal func(int, int) bool
	cal = func(i, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return false
		}
		grid[i][j] = '0'
		cal(i-1, j)
		cal(i, j-1)
		cal(i+1, j)
		cal(i, j+1)
		return true
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if cal(i, j) {
				o++
			}
		}
	}
	return o
}
