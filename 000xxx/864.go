var dirs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type Node struct {
	r, c, nk, step int
}

func shortestPathAllKeys(grid []string) int {
	o, m, n := -1, len(grid), len(grid[0])
	v := make([]bool, 1<<16)
	r, c, nl := -1, -1, uint(0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b := grid[i][j]
			if b == '@' {
				r, c = i, j
			} else if b >= 'a' && b <= 'f' {
				nl++
			}
		}
	}
	if r == -1 {
		return -1
	}

	q := []Node{{r, c, 0, 0}}
	v[(r<<11)|(c<<6)] = true
	for 0 != len(q) {
		t := q[0]
		q = q[1:]
		if t.nk == (1<<nl)-1 {
			o = t.step
			break
		}
		for _, d := range dirs {
			r, c := t.r+d[0], t.c+d[1]
			if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == '#' {
				continue
			}
			b := grid[r][c]
			if b == '#' || (b >= 'A' && b <= 'F' && (t.nk>>(b-'A'))&1 == 0) {
				continue
			}
			nk := t.nk
			if b >= 'a' && b <= 'f' {
				nk |= 1 << (b - 'a')
			}
			k := (r << 11) | (c << 6) | nk
			if !v[k] {
				q = append(q, Node{r, c, nk, t.step + 1})
				v[k] = true
			}
		}
	}
	return o
}
