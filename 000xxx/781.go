func numRabbits(answers []int) int {
	out := 0
	m := map[int]int{}
	for _, a := range answers {
		if 0 == m[a+1] {
			m[a+1] = a
			out += a + 1
		} else {
			m[a+1] = m[a+1] - 1
		}
	}
	return out
}
