func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	dx, dy := xCenter-max(x1, min(xCenter, x2)), yCenter-max(y1, min(yCenter, y2))
	return dx*dx+dy*dy <= radius*radius
}
