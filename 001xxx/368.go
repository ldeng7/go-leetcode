var dirs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	v := make([][]int, m+5)
	for i := 0; i < m+5; i++ {
		ar := make([]int, n+5)
		for j := 0; j < n+5; j++ {
			ar[j] = -1
		}
		v[i] = ar
	}
	ar := make([]uint16, 0, 32)

	var cal func(int, int, int)
	cal = func(i, j, val int) {
		k := grid[i][j] - 1
		r, c := i+dirs[k][0], j+dirs[k][1]
		v[i][j] = val
		if r >= 0 && r < m && c >= 0 && c < n && -1 == v[r][c] {
			cal(r, c, val)
		}
		ar = append(ar, (uint16(i)<<8)|uint16(j))
	}

	cal(0, 0, 0)
	q := make([]uint16, len(ar))
	copy(q, ar)
	ar = make([]uint16, 0, 32)
	for 0 != len(q) {
		e := q[0]
		q = q[1:]
		i, j := int(e>>8), int(e&0xff)
		val := v[i][j]
		for _, dir := range dirs {
			r, c := i+dir[0], j+dir[1]
			if r >= 0 && r < m && c >= 0 && c < n && -1 == v[r][c] {
				cal(r, c, val+1)
				q = append(q, ar...)
				ar = make([]uint16, 0, 32)
			}
		}
	}
	return v[m-1][n-1]
}
