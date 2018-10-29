import "sort"

func hIndex(citations []int) int {
	sort.IntSlice.Sort(citations)
	l := len(citations)
	for i := 0; i < l; i++ {
		if i >= citations[l-i-1] {
			return i
		}
	}
	return l
}
