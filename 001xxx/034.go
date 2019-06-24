func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	v, v1 := map[uint16]bool{}, map[uint16]bool{}
	var cal func(r, c int) uint8
	cal = func(r, c int) uint8 {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] != grid[r0][c0] {
			return 0
		}
		k := (uint16(r) << 8) | uint16(c)
		if v1[k] {
			return 1
		}
		v1[k] = true
		if cal(r+1, c)+cal(r-1, c)+cal(r, c+1)+cal(r, c-1) < 4 {
			v[k] = true
		}
		return 1
	}
	cal(r0, c0)
	for k, _ := range v {
		r, c := k>>8, k&0xff
		grid[r][c] = color
	}
	return grid
}
