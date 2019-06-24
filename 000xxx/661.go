var dirs = [8][2]int{{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}}

func imageSmoother(M [][]int) [][]int {
	if 0 == len(M) || 0 == len(M[0]) {
		return [][]int{}
	}
	m, n := len(M), len(M[0])
	o := make([][]int, m)
	for i := 0; i < m; i++ {
		o[i] = make([]int, n)
		for j := 0; j < n; j++ {
			c, c1 := M[i][j], 1
			for _, d := range dirs {
				y, x := i+d[0], j+d[1]
				if y < 0 || y >= m || x < 0 || x >= n {
					continue
				}
				c, c1 = c+M[y][x], c1+1
			}
			o[i][j] = c / c1
		}
	}
	return o
}
