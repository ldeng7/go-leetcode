func expectNumber(scores []int) int {
	m := map[int]bool{}
	for _, s := range scores {
		m[s] = true
	}
	return len(m)
}
