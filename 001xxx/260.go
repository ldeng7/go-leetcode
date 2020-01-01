func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	l := m * n
	k %= l
	o := make([][]int, m)
	for i := 0; i < m; i++ {
		o[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p := (l - k + i*n + j) % l
			o[i][j] = grid[p/n][p%n]
		}
	}
	return o
}
