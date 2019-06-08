import "sort"

func findRightInterval(intervals [][]int) []int {
	out, t, m := []int{}, []int{}, map[int]int{}
	for i, itvl := range intervals {
		m[itvl[0]], t = i, append(t, itvl[0])
	}
	sort.Ints(t)
	for _, itvl := range intervals {
		j := sort.Search(len(t), func(j int) bool {
			return t[j] >= itvl[1]
		})
		if j != len(t) {
			out = append(out, m[t[j]])
		} else {
			out = append(out, -1)
		}
	}
	return out
}
