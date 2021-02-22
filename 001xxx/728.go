var dirs = [5][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}, {0, 0}}

type Node struct {
	my, mx, cy, cx, p int
}

func canMouseWin(grid []string, catJump int, mouseJump int) bool {
	m, n := len(grid), len(grid[0])
	t, s := [8][8][8][8][2]int{}, [8][8][8][8][2]int{}
	c := [8][8][2]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '#' {
				continue
			}
			c[i][j][0], c[i][j][1] = 1, 1
			for h := 0; h < 4; h++ {
				d := dirs[h]
				for k := 1; k <= mouseJump; k++ {
					y, x := i+d[0]*k, j+d[1]*k
					if y < 0 || y >= m || x < 0 || x >= n || grid[y][x] == '#' {
						break
					}
					c[i][j][0]++
				}
				for k := 1; k <= catJump; k++ {
					y, x := i+d[0]*k, j+d[1]*k
					if y < 0 || y >= m || x < 0 || x >= n || grid[y][x] == '#' {
						break
					}
					c[i][j][1]++
				}
			}
		}
	}

	var my0, mx0, fy0, fx0, cy0, cx0 int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case 'M':
				my0, mx0 = i, j
			case 'F':
				fy0, fx0 = i, j
			case 'C':
				cy0, cx0 = i, j
			}
			for k := 0; k < m; k++ {
				for h := 0; h < n; h++ {
					s[i][j][k][h] = [2]int{c[i][j][0], c[k][h][1]}
				}
			}
		}
	}

	q := make([]Node, 0, 16)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '#' || (i == fy0 && j == fx0) {
				continue
			}
			q = append(q,
				Node{i, j, fy0, fx0, 0}, Node{i, j, fy0, fx0, 1},
				Node{fy0, fx0, i, j, 0}, Node{fy0, fx0, i, j, 1},
				Node{i, j, i, j, 0}, Node{i, j, i, j, 1})
			t[i][j][fy0][fx0] = [2]int{2, 2}
			t[fy0][fx0][i][j] = [2]int{1, 1}
			t[i][j][i][j] = [2]int{2, 2}
		}
	}

	for 0 != len(q) {
		node := q[0]
		q = q[1:]
		my, mx, cy, cx, p := node.my, node.mx, node.cy, node.cx, node.p
		if my == my0 && mx == mx0 && cy == cy0 && cx == cx0 && p == 0 {
			break
		}
		st := t[my][mx][cy][cx][p]
		if p == 1 {
			for i, d := range dirs {
				for j := 1; j <= mouseJump; j++ {
					y, x := my+d[0]*j, mx+d[1]*j
					if y < 0 || y >= m || x < 0 || x >= n || grid[y][x] == '#' {
						break
					} else if 0 != t[y][x][cy][cx][p^1] {
						continue
					}
					st1 := 0
					if st == 1 {
						st1 = 1
					} else {
						v := &(s[y][x][cy][cx][p^1])
						(*v)--
						if 0 == *v {
							st1 = 2
						}
					}
					if 0 != st1 {
						t[y][x][cy][cx][p^1] = st1
						q = append(q, Node{y, x, cy, cx, p ^ 1})
					}
					if i == 4 {
						break
					}
				}
			}
		} else {
			for i, d := range dirs {
				for j := 1; j <= catJump; j++ {
					y, x := cy+d[0]*j, cx+d[1]*j
					if y < 0 || y >= m || x < 0 || x >= n || grid[y][x] == '#' {
						break
					} else if 0 != t[my][mx][y][x][p^1] {
						continue
					}
					st1 := 0
					if st == 2 {
						st1 = 2
					} else {
						v := &(s[my][mx][y][x][p^1])
						(*v)--
						if 0 == *v {
							st1 = 1
						}
					}
					if 0 != st1 {
						t[my][mx][y][x][p^1] = st1
						q = append(q, Node{my, mx, y, x, p ^ 1})
					}
					if i == 4 {
						break
					}
				}
			}
		}
	}
	return t[my0][mx0][cy0][cx0][0] == 1
}
