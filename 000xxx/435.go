import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	o, e := 0, 0
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	for i := 1; i < len(intervals); i++ {
		a, b := intervals[i], intervals[e]
		if a[0] < b[1] {
			o++
			if a[1] < b[1] {
				e = i
			}
		} else {
			e = i
		}
	}
	return o
}
