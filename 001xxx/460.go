func canBeEqual(target []int, arr []int) bool {
	m1, m2 := map[int]int{}, map[int]int{}
	for _, n := range target {
		m1[n]++
	}
	for _, n := range arr {
		m2[n]++
	}
	if len(m1) != len(m2) {
		return false
	}
	for n, c := range m1 {
		if c != m2[n] {
			return false
		}
	}
	return true
}
