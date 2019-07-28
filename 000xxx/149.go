func maxPoints(points [][]int) int {
	o, l := 0, len(points)
	for i := 0; i < l; i++ {
		d := 1
		for j := i + 1; j < l; j++ {
			c := 0
			y1, x1, y2, x2 := points[i][0], points[i][1], points[j][0], points[j][1]
			if x1 == x2 && y1 == y2 {
				d++
				continue
			}
			for k := 0; k < l; k++ {
				y3, x3 := points[k][0], points[k][1]
				if x1*y2+x2*y3+x3*y1-x3*y2-x2*y1-x1*y3 == 0 {
					c++
				}
			}
			if c > o {
				o = c
			}
		}
		if d > o {
			o = d
		}
	}
	return o
}
