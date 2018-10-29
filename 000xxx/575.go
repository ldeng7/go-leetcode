func distributeCandies(candies []int) int {
	m := map[int]bool{}
	for _, c := range candies {
		m[c] = true
	}
	a, b := len(m), len(candies)/2
	if a < b {
		return a
	}
	return b
}
