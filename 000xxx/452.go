import "sort"

type Pairs [][]int

func (p Pairs) Len() int { return len(p) }
func (p Pairs) Less(i, j int) bool {
	return p[i][0] < p[j][0] || (p[i][0] == p[j][0] && p[i][1] < p[j][1])
}
func (p Pairs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func findMinArrowShots(points [][]int) int {
	if 0 == len(points) {
		return 0
	}
	sort.Sort(Pairs(points))
	o, e := 1, points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= e {
			e = min(e, points[i][1])
		} else {
			o++
			e = points[i][1]
		}
	}
	return o
}
