func anagramMappings(A []int, B []int) []int {
	m := map[int]int{}
	for i, n := range B {
		m[n] = i
	}
	out := make([]int, len(A))
	for i, n := range A {
		out[i] = m[n]
	}
	return out
}
