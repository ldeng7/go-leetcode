func maxPoints(points [][]int) int {
	out := 0
	for i := 0; i < len(points); i++ {
		d := 1
		for j := i + 1; j < len(points); j++ {
			c := 0
			y1, x1, y2, x2 := points[i][0], points[i][1], points[j][0], points[j][1]
			if x1 == x2 && y1 == y2 {
				d++
				continue
			}
			for k := 0; k < len(points); k++ {
				y3, x3 := points[k][0], points[k][1]
				if x1*y2+x2*y3+x3*y1-x3*y2-x2*y1-x1*y3 == 0 {
					c++
				}
			}
			if c > out {
				out = c
			}
		}
		if d > out {
			out = d
		}
	}
	return out
}
