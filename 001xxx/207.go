func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}
	for _, a := range arr {
		m[a]++
	}
	m1 := map[int]bool{}
	for _, c := range m {
		if m1[c] {
			return false
		}
		m1[c] = true
	}
	return true
}
