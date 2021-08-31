var dirs = [5][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {0, 0}}

func escapeMaze(maze [][]string) bool {
	l, m, n := len(maze)-1, len(maze[0]), len(maze[0][0])
	v := [100][50][50][4]bool{}
	var cal func(int, int, int, bool, bool) bool
	cal = func(y, x, s int, b1, b2 bool) bool {
		i := 0
		if b1 {
			i |= 2
		}
		if b2 {
			i |= 1
		}
		if v[s][y][x][i] {
			return false
		}
		v[s][y][x][i] = true
		if y == m-1 && x == n-1 {
			return true
		} else if s == l {
			return false
		} else if l-s < m-1-y+n-1-x {
			return false
		}
		for _, dir := range dirs {
			y1, x1 := y+dir[0], x+dir[1]
			if y1 < 0 || y1 >= m || x1 < 0 || x1 >= n {
				continue
			}
			if maze[s+1][y1][x1] == '.' {
				if cal(y1, x1, s+1, b1, b2) {
					return true
				}
				continue
			}
			if !b1 {
				if cal(y1, x1, s+1, true, b2) {
					return true
				}
			}
			if !b2 {
				for i := s + 1; i <= l; i++ {
					if cal(y1, x1, i, b1, true) {
						return true
					}
				}
			}
		}
		return false
	}
	return cal(0, 0, 0, false, false)
}
