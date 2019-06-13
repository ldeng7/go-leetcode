import "sort"

func intersectionSizeTwo(intervals [][]int) int {
	o, p1, p2 := 0, -1, -1
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[1] < b[1] || (a[1] == b[1] && a[0] > b[0])
	})
	for _, itvl := range intervals {
		if itvl[0] <= p1 {
			continue
		}
		if itvl[0] > p2 {
			o, p1, p2 = o+2, itvl[1]-1, itvl[1]
		} else {
			o, p1, p2 = o+1, p2, itvl[1]
		}
	}
	return o
}
