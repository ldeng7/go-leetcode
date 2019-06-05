func numDistinctIslands(grid [][]int) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	m, n := len(grid), len(grid[0])
	var cal func(int, int, int, int, *[]byte)
	cal = func(i, j, i0, j0 int, bs *[]byte) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		*bs = append(*bs, byte(i0)+50-byte(i))
		*bs = append(*bs, byte(j0)+50-byte(j))
		cal(i-1, j, i0, j0, bs)
		cal(i, j-1, i0, j0, bs)
		cal(i+1, j, i0, j0, bs)
		cal(i, j+1, i0, j0, bs)
	}
	ms := map[string]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			bs := &[]byte{}
			cal(i, j, i, j, bs)
			if 0 != len(*bs) {
				ms[string(*bs)] = true
			}
		}
	}
	return len(ms)
}
