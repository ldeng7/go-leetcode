import "math"

func dist(x0, y0, x1, y1 float64) float64 {
	dx, dy := x1-x0, y1-y0
	return dx*dx + dy*dy
}

func numPoints(points [][]int, r int) int {
	o, l := 1, len(points)
	rr := float64(r * r)
	for i, p0 := range points {
		x0, y0 := float64(p0[0]), float64(p0[1])
		for j := i + 1; j < l; j++ {
			p1 := points[j]
			x1, y1 := float64(p1[0]), float64(p1[1])
			d := dist(x0, y0, x1, y1)
			if d > rr*4 {
				continue
			}
			xm, ym := (x0+x1)/2, (y0+y1)/2
			s, sd := math.Sqrt(rr-dist(x0, y0, xm, ym)), math.Sqrt(d)
			x, y := xm-s*(y1-y0)/sd, ym+s*(x1-x0)/sd
			c := 0
			for k := 0; k < l; k++ {
				if k == i || k == j {
					c++
					continue
				}
				p := points[k]
				if dist(float64(p[0]), float64(p[1]), x, y) <= rr {
					c++
				}
			}
			if c > o {
				o = c
			}
		}
	}
	return o
}
