import "math"

func dist(ya, xa, yb, xb int) float64 {
	return math.Sqrt(float64((yb-ya)*(yb-ya) + (xb-xa)*(xb-xa)))
}

func minAreaFreeRect(points [][]int) float64 {
	l := len(points)
	v := map[int]map[int][][2]int{}
	for i := 0; i < l; i++ {
		pa := [2]int{points[i][0], points[i][1]}
		for j := i + 1; j < l; j++ {
			pb := points[j]
			cy, cx := pa[0]+pb[0], pa[1]+pb[1]
			r := (pa[0]-pb[0])*(pa[0]-pb[0]) + (pa[1]-pb[1])*(pa[1]-pb[1])
			if _, ok := v[r]; !ok {
				v[r] = map[int][][2]int{}
			}
			k, m := (cy<<16)|cx, v[r]
			m[k] = append(m[k], pa)
		}
	}
	o := math.MaxFloat64
	for _, m := range v {
		for k, ar := range m {
			cy, cx := k>>16, k&0xffff
			for i, pa := range ar {
				for j := i + 1; j < len(ar); j++ {
					pb := ar[j]
					a := dist(pa[0], pa[1], pb[0], pb[1]) * dist(pa[0], pa[1], cy-pb[0], cx-pb[1])
					if a < o {
						o = a
					}
				}
			}
		}
	}
	if o == math.MaxFloat64 {
		return 0
	}
	return o
}
