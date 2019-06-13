import "sort"

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
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || (points[i][0] == points[j][0] && points[i][1] < points[j][1])
	})
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
