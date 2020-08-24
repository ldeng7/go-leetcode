import "math"

func findMaxValueOfEquation(points [][]int, k int) int {
	q := make([]int, 0, len(points))
	o := math.MinInt32
	for i, p := range points {
		x, y := p[0], p[1]
		for 0 != len(q) && x-points[q[0]][0] > k {
			q = q[1:]
		}
		if 0 != len(q) {
			p1 := points[q[0]]
			if t := x + y - (p1[0] - p1[1]); t > o {
				o = t
			}
		}
		for 0 != len(q) {
			p1 := points[q[len(q)-1]]
			if x-y > p1[0]-p1[1] {
				break
			}
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return o
}
