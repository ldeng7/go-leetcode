import "sort"

func getKey(ars []*[]int) string {
	ss := make([]string, 8)
	for k := 0; k < 8; k++ {
		ar := *ars[k]
		sort.Ints(ar)
		y0, x0 := byte(ar[0]>>8), byte(ar[0]&0xff)
		bs := make([]byte, len(ar)<<1)
		for i, e := range ar {
			bs[i<<1], bs[i<<1+1] = byte(e>>8)+50-y0, byte(e&0xff)+50-x0
		}
		ss[k] = string(bs)
	}
	sort.Strings(ss)
	return ss[0]
}

func numDistinctIslands2(grid [][]int) int {
	if 0 == len(grid) || 0 == len(grid[0]) {
		return 0
	}
	m, n := len(grid), len(grid[0])
	var cal func(int, int, int, int, []*[]int)
	cal = func(i, j, i0, j0 int, ars []*[]int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		p2i := func(di, dj int) int {
			return ((di + 50) << 8) | (dj + 50)
		}
		di, dj := i-i0, j-j0
		*ars[0] = append(*ars[0], p2i(di, dj))
		*ars[1] = append(*ars[1], p2i(-di, dj))
		*ars[2] = append(*ars[2], p2i(di, -dj))
		*ars[3] = append(*ars[3], p2i(-di, -dj))
		*ars[4] = append(*ars[4], p2i(dj, di))
		*ars[5] = append(*ars[5], p2i(-dj, di))
		*ars[6] = append(*ars[6], p2i(dj, -di))
		*ars[7] = append(*ars[7], p2i(-dj, -di))
		cal(i-1, j, i0, j0, ars)
		cal(i, j-1, i0, j0, ars)
		cal(i+1, j, i0, j0, ars)
		cal(i, j+1, i0, j0, ars)
	}
	ms := map[string]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ars := make([]*[]int, 8)
			for k := 0; k < 8; k++ {
				ars[k] = &[]int{}
			}
			cal(i, j, i, j, ars)
			if 0 != len(*ars[0]) {
				ms[getKey(ars)] = true
			}
		}
	}
	return len(ms)
}
