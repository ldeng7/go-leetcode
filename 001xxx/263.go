var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minPushBox(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	l := m * n
	var s, b, t int
	for i, line := range grid {
		for j, p := range line {
			switch p {
			case 'S':
				s = i*n + j
			case 'B':
				b = i*n + j
			case 'T':
				t = i*n + j
			}
		}
	}

	check := func(s, b, t int) bool {
		if s == t {
			return true
		}
		v := make([]bool, l)
		q := make([]int, 1, 16)
		q[0] = s
		for 0 != len(q) {
			p := q[0]
			q = q[1:]
			r, c := p/n, p%n
			for _, d := range dirs {
				y, x := r+d[0], c+d[1]
				p = y*n + x
				if y >= 0 && y < m && x >= 0 && x < n && grid[y][x] != '#' && !v[p] && p != b {
					if p == t {
						return true
					}
					q = append(q, p)
					v[p] = true
				}
			}
		}
		return false
	}

	v := make([]bool, l<<2)
	q := make([]uint32, 1, 16)
	q[0] = (uint32(b) << 16) | uint32(s)
	o, lq := 1, len(q)
	for 0 != len(q) {
		for 0 != lq {
			p := q[0]
			q, lq = q[1:], lq-1
			b1, s1 := int(p>>16), int(p&0xffff)
			r, c := b1/n, b1%n
			for i, d := range dirs {
				y, x, ys, xs := r+d[0], c+d[1], r-d[0], c-d[1]
				p, s2 := y*n+x, ys*n+xs
				pv := p<<2 | i
				if y >= 0 && y < m && x >= 0 && x < n && ys >= 0 && ys < m && xs >= 0 && xs < n &&
					grid[y][x] != '#' && grid[ys][xs] != '#' && !v[pv] && check(s1, b1, s2) {
					if p == t {
						return o
					}
					q = append(q, (uint32(p)<<16)|uint32(s2))
					v[pv] = true
				}
			}
			if 0 == lq {
				o, lq = o+1, len(q)
			}
		}
	}
	return -1
}
