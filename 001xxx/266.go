func minTimeToVisitAllPoints(points [][]int) int {
	o, l := 0, len(points)
	xp, yp := points[0][0], points[0][1]
	for i := 1; i < l; i++ {
		x, y := points[i][0], points[i][1]
		dx, dy := x-xp, y-yp
		if dx < 0 {
			dx = -dx
		}
		if dy < 0 {
			dy = -dy
		}
		if dx >= dy {
			o += dx
		} else {
			o += dy
		}
		xp, yp = x, y
	}
	return o
}
