var dirs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	o := make([][]int, m)
	for i := 0; i < m; i++ {
		ar := make([]int, n)
		for j := 0; j < n; j++ {
			ar[j] = -1
		}
		o[i] = ar
	}

	q := make([]uint32, 0, 16)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				o[i][j] = 0
				q = append(q, uint32(i<<16)|uint32(j))
			}
		}
	}

	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		y, x := int(t>>16), int(t&0xffff)
		d := o[y][x]
		for _, dir := range dirs {
			y1, x1 := y+dir[0], x+dir[1]
			if y1 >= 0 && y1 < m && x1 >= 0 && x1 < n && o[y1][x1] == -1 {
				o[y1][x1] = d + 1
				q = append(q, uint32(y1<<16)|uint32(x1))
			}
		}
	}
	return o
}
