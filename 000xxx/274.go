import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	l := len(citations)
	for i := 0; i < l; i++ {
		if i >= citations[l-i-1] {
			return i
		}
	}
	return l
}
