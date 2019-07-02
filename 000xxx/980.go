func uniquePathsIII(grid [][]int) int {
	v := map[int]bool{}
	m, n := len(grid), len(grid[0])
	o, nc, r, c := 0, 1, 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case 0:
				nc++
			case 1:
				r, c = i, j
			}
		}
	}

	var cal func(int, int, int)
	cal = func(i, j, nc int) {
		if grid[i][j] == 2 {
			if nc == 0 {
				o++
			}
			return
		}
		k := (i << 8) | j
		v[k] = true
		if i-1 >= 0 && grid[i-1][j] != -1 && !v[k-256] {
			cal(i-1, j, nc-1)
		}
		if j-1 >= 0 && grid[i][j-1] != -1 && !v[k-1] {
			cal(i, j-1, nc-1)
		}
		if i+1 < m && grid[i+1][j] != -1 && !v[k+256] {
			cal(i+1, j, nc-1)
		}
		if j+1 < n && grid[i][j+1] != -1 && !v[k+1] {
			cal(i, j+1, nc-1)
		}
		v[k] = false
	}
	cal(r, c, nc)
	return o
}
