func groupThePeople(groupSizes []int) [][]int {
	m := map[int][]int{}
	o := [][]int{}
	for i, sz := range groupSizes {
		ar := append(m[sz], i)
		if len(ar) == sz {
			o = append(o, ar)
			m[sz] = []int{}
		} else {
			m[sz] = ar
		}
	}
	return o
}
