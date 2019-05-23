func maxAreaOfIsland(grid [][]int) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	m, n := len(grid), len(grid[0])
	out := 0
	var cal func(int, int) int
	cal = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}
		out := 1
		grid[i][j] = 0
		out += cal(i-1, j)
		out += cal(i, j-1)
		out += cal(i+1, j)
		out += cal(i, j+1)
		return out
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a := cal(i, j)
			if a > out {
				out = a
			}
		}
	}
	return out
}
