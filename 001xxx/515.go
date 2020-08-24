import "math"

var dirs = [4][2]float64{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func getMinDistSum(positions [][]int) float64 {
	dist := func(x, y float64) float64 {
		var d float64
		for _, p := range positions {
			dx, dy := x-float64(p[0]), y-float64(p[1])
			d += math.Sqrt(dx*dx + dy*dy)
		}
		return d
	}

	l := float64(len(positions))
	var x, y float64
	for _, p := range positions {
		x, y = x+float64(p[0]), y+float64(p[1])
	}
	x, y = x/l, y/l
	var d, s float64 = dist(x, y), 50
	ns := 0

	for s > 0.00000001 && ns < 1000000 {
		ns++
		b := false
		for _, dir := range dirs {
			x1, y1 := x+dir[0]*s, y+dir[1]*s
			d1 := dist(x1, y1)
			if d1 < d {
				x, y, d, b = x1, y1, d1, true
				break
			}
		}
		if !b {
			s /= 2
		}
	}
	return d
}
