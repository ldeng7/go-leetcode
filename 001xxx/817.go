func findingUsersActiveMinutes(logs [][]int, k int) []int {
	m := map[int]map[int]bool{}
	for _, l := range logs {
		i := l[0]
		if m1, ok := m[i]; ok {
			m1[l[1]] = true
		} else {
			m[i] = map[int]bool{l[1]: true}
		}
	}
	o := make([]int, k)
	for _, m1 := range m {
		o[len(m1)-1]++
	}
	return o
}
