func visitOrder(points [][]int, direction string) []int {
	l := len(points)
	o := make([]int, l)
	for i := 0; i < l; i++ {
		o[i] = i
	}

	var cal func(int) bool
	cal = func(k int) bool {
		if k == l {
			return true
		}
		var dir byte
		var x1, y1, x2, y2 int
		if k >= 2 {
			p1, p2 := points[o[k-2]], points[o[k-1]]
			x1, y1, x2, y2 = p1[0], p1[1], p2[0], p2[1]
			dir = direction[k-2]
		}
		for i := k; i < l; i++ {
			if k >= 2 {
				p := points[o[i]]
				v := (y1-y2)*(p[0]-x2) + (x2-x1)*(p[1]-y2)
				if (dir == 'L' && v < 0) || (dir == 'R' && v > 0) {
					continue
				}
			}
			a, b := o[k], o[i]
			o[k], o[i] = b, a
			if cal(k + 1) {
				return true
			}
			o[k], o[i] = a, b
		}
		return false
	}

	cal(0)
	return o
}
