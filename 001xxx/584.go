func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func minCostConnectPoints(points [][]int) int {
	o, l := 0, len(points)
	h, d := make([]bool, l), make([]int, l)
	for i := 0; i < l; i++ {
		d[i] = 1e7
	}
	for i, n := 0, l; n > 1; n-- {
		x, y, k := points[i][0], points[i][1], -1
		h[i] = true
		for j := l - 1; j >= 0; j-- {
			if h[j] {
				continue
			}
			t := abs(x-points[j][0]) + abs(y-points[j][1])
			if d[j] > t {
				d[j] = t
			}
			if k == -1 || d[k] > d[j] {
				k = j
			}
		}
		o += d[k]
		i = k
	}
	return o
}
