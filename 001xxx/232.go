func checkStraightLine(coordinates [][]int) bool {
	l := len(coordinates)
	if l <= 2 {
		return true
	}
	var x0, y0 int
	x1, y1 := coordinates[0][0], coordinates[0][1]
	x2, y2 := coordinates[1][0], coordinates[1][1]
	d := x1*y2 - x2*y1
	for i := 2; i < l; i++ {
		x0, y0, x1, y1, x2, y2 = x1, y1, x2, y2, coordinates[i][0], coordinates[i][1]
		pd := d
		d = x1*y2 - x2*y1
		if s := pd + x2*y0 - x0*y2 + d; s != 0 {
			return false
		}
	}
	return true
}
