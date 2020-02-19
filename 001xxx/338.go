import "sort"

func minSetSize(arr []int) int {
	m := map[int]int{}
	for _, a := range arr {
		m[a]++
	}
	cs := make([]int, 0, len(m))
	for _, c := range m {
		cs = append(cs, c)
	}
	sort.Slice(cs, func(i, j int) bool { return cs[i] > cs[j] })

	s, l := 0, len(arr)
	for i := 0; i < l; i++ {
		s += cs[i]
		if s >= l>>1 {
			return i + 1
		}
	}
	return l
}
