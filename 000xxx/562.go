func longestLine(M [][]int) int {
	if 0 == len(M) || 0 == len(M[0]) {
		return 0
	}
	m, n := len(M), len(M[0])
	out := 0
	dirs := [][]int{{1, 0}, {0, 1}, {-1, -1}, {-1, 1}}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if 0 == M[i][j] {
				continue
			}
			for _, d := range dirs {
				c, y, x := 0, i, j
				for y >= 0 && y < m && x >= 0 && x < n && M[y][x] == 1 {
					c, y, x = c+1, y+d[0], x+d[1]
				}
				if c > out {
					out = c
				}
			}
		}
	}
	return out
}
