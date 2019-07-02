var dirs = [8][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

func shortestPathBinaryMatrix(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if grid[0][0] == 1 || grid[m-1][n-1] == 1 {
		return -1
	}
	grid[0][0] = -1
	q := [][3]int{{0, 0, 1}}
	for 0 != len(q) {
		l := len(q)
		for i := 0; i < l; i++ {
			t := q[0]
			q = q[1:]
			y0, x0, v0 := t[0], t[1], t[2]
			if y0 == m-1 && x0 == n-1 {
				return v0
			}
			for _, d := range dirs {
				y, x := y0+d[0], x0+d[1]
				if y >= 0 && y < m && x >= 0 && x < n && grid[y][x] == 0 {
					grid[y][x] = -1
					q = append(q, [3]int{y, x, v0 + 1})
				}
			}
		}
	}
	return -1
}
