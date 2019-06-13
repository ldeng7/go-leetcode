func hitBricks(grid [][]int, hits [][]int) []int {
	m, n, l := len(grid), len(grid[0]), len(hits)
	out := make([]int, l)
	s := map[int]bool{}
	for i := 0; i < l; i++ {
		grid[hits[i][0]][hits[i][1]] -= 1
	}

	var cal func(i, j int)
	cal = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 || s[i*n+j] {
			return
		}
		s[i*n+j] = true
		cal(i-1, j)
		cal(i+1, j)
		cal(i, j-1)
		cal(i, j+1)
	}

	for i := 0; i < n; i++ {
		if grid[0][i] == 1 {
			cal(0, i)
		}
	}
	for k := l - 1; k >= 0; k-- {
		ls, i, j := len(s), hits[k][0], hits[k][1]
		grid[i][j]++
		if grid[i][j] != 1 {
			continue
		}
		if (i-1 >= 0 && s[(i-1)*n+j]) || (i+1 < m && s[(i+1)*n+j]) ||
			(j-1 >= 0 && s[i*n+j-1]) || (j+1 < n && s[i*n+j+1]) || i == 0 {
			cal(i, j)
			out[k] = len(s) - ls - 1
		}
	}
	return out
}
