import "sort"

func countPoints(points [][]int, queries [][]int) []int {
	ar := make([]int, len(points))
	for i, p := range points {
		ar[i] = (p[0] << 16) | p[1]
	}
	sort.Ints(ar)
	o := make([]int, len(queries))
	for i, q := range queries {
		c, x0, y0, r := 0, q[0], q[1], q[2]
		rr, p := r*r, (x0-r)<<16
		for j := sort.Search(len(ar), func(k int) bool {
			return ar[k] >= p
		}); j != len(ar); j++ {
			p := ar[j]
			x, y := p>>16, p&0xffff
			if x > x0+r {
				break
			}
			x, y = x0-x, y0-y
			if x*x+y*y <= rr {
				c++
			}
		}
		o[i] = c
	}
	return o
}
