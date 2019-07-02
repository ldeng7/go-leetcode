var dirs = [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	bm := map[int]bool{}
	for _, b := range blocked {
		bm[b[0]<<32+b[1]] = true
	}
	cal := func(sy, sx, ty, tx int) bool {
		q := make([][2]int, 1, 64)
		v := map[int]bool{(sy << 32) | sx: true}
		q[0] = [2]int{sy, sx}
		c := 0
		for len(q) > 0 {
			y0, x0 := q[0][0], q[0][1]
			for _, d := range dirs {
				y, x := y0+d[0], x0+d[1]
				if y == ty && x == tx {
					return true
				}
				k := (y << 32) | x
				if x >= 0 && x < 1000000 && y >= 0 && y < 1000000 && !bm[k] && !v[k] {
					q, v[k] = append(q, [2]int{y, x}), true
				}
			}
			q, c = q[1:], c+1
			if c > 19900 {
				return true
			}
		}
		return false
	}
	return cal(source[0], source[1], target[0], target[1]) &&
		cal(target[0], target[1], source[0], source[1])
}
