func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	o := 0
	var cal func(int, int) int
	cal = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return -1
		} else if grid[i][j] == 1 {
			return 0
		}
		grid[i][j] = 1
		o1, o2, o3, o4 := cal(i-1, j), cal(i, j-1), cal(i+1, j), cal(i, j+1)
		if o1 == -1 || o2 == -1 || o3 == -1 || o4 == -1 {
			return -1
		}
		return 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if cal(i, j) == 1 {
				o++
			}
		}
	}
	return o
}
