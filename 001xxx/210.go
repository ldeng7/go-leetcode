type node struct {
	y, x byte
	dir  bool
	l    uint16
}

func minimumMoves(grid [][]int) int {
	o, n := -1, byte(len(grid))
	v := make([][][2]bool, n)
	for i := byte(0); i < n; i++ {
		v[i] = make([][2]bool, n)
	}
	v[0][0][0] = true
	q := []node{{0, 0, false, 0}}
	for 0 != len(q) {
		h := q[0]
		q = q[1:]
		y, x, dir, l := h.y, h.x, h.dir, h.l

		if y == n-1 && x == n-2 && !dir {
			o = int(l)
			break
		} else if !dir {
			if x+2 < n && grid[y][x+2] == 0 && !v[y][x+1][0] {
				v[y][x+1][0] = true
				q = append(q, node{y, x + 1, false, l + 1})
			}
			if y+1 < n && grid[y+1][x] == 0 && grid[y+1][x+1] == 0 {
				if !v[y+1][x][0] {
					v[y+1][x][0] = true
					q = append(q, node{y + 1, x, false, l + 1})
				}
				if !v[y][x][1] {
					v[y][x][1] = true
					q = append(q, node{y, x, true, l + 1})
				}
			}
		} else {
			if y+2 < n && grid[y+2][x] == 0 && !v[y+1][x][1] {
				v[y+1][x][1] = true
				q = append(q, node{y + 1, x, true, l + 1})
			}
			if x+1 < n && grid[y][x+1] == 0 && grid[y+1][x+1] == 0 {
				if !v[y][x+1][1] {
					v[y][x+1][1] = true
					q = append(q, node{y, x + 1, true, l + 1})
				}
				if !v[y][x][0] {
					v[y][x][0] = true
					q = append(q, node{y, x, false, l + 1})
				}
			}
		}
	}
	return o
}
