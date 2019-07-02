var dirs = [4][2]int8{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func orangesRotting(grid [][]int) int {
	o, m, n := 0, uint8(len(grid)), uint8(len(grid[0]))
	q, md := []uint8{}, map[uint8]int{}
	for i := uint8(0); i < m; i++ {
		for j := uint8(0); j < n; j++ {
			if grid[i][j] == 2 {
				k := (i << 4) | j
				q, md[k] = append(q, k), 0
			}
		}
	}
	for 0 != len(q) {
		k := q[0]
		q = q[1:]
		var y, x uint8 = k >> 4, k & 0x0f
		for _, d := range dirs {
			i, j := int8(y)+d[0], int8(x)+d[1]
			if i < 0 || i >= int8(m) || j < 0 || j >= int8(n) || grid[i][j] != 1 {
				continue
			}
			grid[i][j] = 2
			k1 := uint8((i << 4) | j)
			q, md[k1] = append(q, k1), md[k]+1
			o = md[k1]
		}
	}

	for i := uint8(0); i < m; i++ {
		for j := uint8(0); j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return o
}
