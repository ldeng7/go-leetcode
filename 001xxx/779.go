func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func nearestValidPoint(x int, y int, points [][]int) int {
	var o, d int = -1, 1e5
	for i, p := range points {
		dx, dy := x-p[0], y-p[1]
		if dx == 0 || dy == 0 {
			if d1 := abs(dx) + abs(dy); d1 < d {
				o, d = i, d1
			}
		}
	}
	return o
}
