import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	l := len(intervals)
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		a0, b0 := a[0], b[0]
		return a0 < b0 || (a0 == b0 && a[1] > b[1])
	})
	o, ar := l, intervals[0]
	for i := 1; i < l; i++ {
		ar1 := intervals[i]
		if ar1[0] >= ar[0] && ar1[1] <= ar[1] {
			o--
		} else {
			ar = ar1
		}
	}
	return o
}
