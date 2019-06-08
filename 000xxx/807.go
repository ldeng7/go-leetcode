func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	o, m, n := 0, len(grid), len(grid[0])
	rs, cs := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			t := grid[i][j]
			rs[i] = max(rs[i], t)
			cs[j] = max(cs[j], t)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			o += min(rs[i]-grid[i][j], cs[j]-grid[i][j])
		}
	}
	return o
}
