var dirs = [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func getMaxGridHappiness(m int, n int, introvertsCount int, extrovertsCount int) int {
	o := 0
	g := make([][]int, n)
	for j := 0; j < n; j++ {
		g[j] = make([]int, m)
	}

	var calIn func(int, int, int)
	calIn = func(i, p, s int) {
		if s+i*120 <= o {
			return
		} else if i == 0 || p == n*m {
			o = max(o, s)
			return
		}
		for j := p; j < n*m; j++ {
			x, y := j/m, j%m
			if g[x][y] != 0 {
				continue
			}
			s1 := 120
			for _, d := range dirs {
				x1, y1 := x+d[0], y+d[1]
				if y1 < 0 || y1 >= m || x1 < 0 || x1 >= n || g[x1][y1] == 0 {
					continue
				}
				s1 -= 30
				if g[x1][y1] == 2 {
					s1 += 20
				} else {
					s1 -= 30
				}
			}
			g[x][y] = 1
			calIn(i-1, j+1, s+s1)
			g[x][y] = 0
		}
	}

	var calEx func(int, int, int)
	calEx = func(i, p, s int) {
		if i == 0 || p == n*m {
			calIn(introvertsCount, 0, s)
			return
		}
		for j := p; j < n*m; j++ {
			x, y := j/m, j%m
			s1 := 40
			for _, d := range dirs {
				x1, y1 := x+d[0], y+d[1]
				if y1 < 0 || y1 >= m || x1 < 0 || x1 >= n || g[x1][y1] == 0 {
					continue
				}
				s1 += 40
			}
			if i != extrovertsCount && s1 == 40 {
				continue
			}
			g[x][y] = 2
			calEx(i-1, j+1, s+s1)
			calEx(i-1, n*m, s+s1)
			g[x][y] = 0
		}
	}

	calEx(extrovertsCount, 0, 0)
	calIn(introvertsCount, 0, 0)
	return o
}
