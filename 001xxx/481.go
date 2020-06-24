import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	m := map[int]int{}
	for _, a := range arr {
		m[a]++
	}
	cs := make([]int, 0, len(m))
	for _, c := range m {
		cs = append(cs, c)
	}
	sort.Ints(cs)

	l := len(cs)
	for i, c := range cs {
		if k < c {
			return l - i
		}
		k -= c
	}
	return 0
}
