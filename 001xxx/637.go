import "sort"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxWidthOfVerticalArea(points [][]int) int {
	o, l := 0, len(points)
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	for i := 1; i < l; i++ {
		o = max(o, points[i][0]-points[i-1][0])
	}
	return o
}
