var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func domino(n int, m int, broken [][]int) int {
	g := make([][]bool, n)
	for i := 0; i < n; i++ {
		g[i] = make([]bool, m)
	}
	for _, b := range broken {
		g[b[0]][b[1]] = true
	}
	l := m * n
	lines := make([][]int, l+1)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if g[i][j] || (i+j)&1 == 1 {
				continue
			}
			idx := i*m + j + 1
			for _, d := range dirs {
				x, y := i+d[0], j+d[1]
				if x < 0 || y < 0 || x >= n || y >= m || g[x][y] {
					continue
				}
				lines[idx] = append(lines[idx], x*m+y+1)
			}
		}
	}

	var v []bool
	ar := make([]int, l+1)
	var cal func(int) bool
	cal = func(i int) bool {
		for _, j := range lines[i] {
			if v[j] {
				continue
			}
			v[j] = true
			if k := ar[j]; k == 0 || cal(k) {
				ar[j] = i
				return true
			}
		}
		return false
	}

	o := 0
	for i := 0; i < l; i++ {
		v = make([]bool, l+1)
		if cal(i + 1) {
			o++
		}
	}
	return o
}
