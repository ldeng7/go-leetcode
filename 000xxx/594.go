func findLHS(nums []int) int {
	out := 0
	m := map[int]int{}
	for _, n := range nums {
		m[n] = m[n] + 1
	}
	for n, c := range m {
		if cn, ok := m[n+1]; ok && c+cn > out {
			out = c + cn
		}
	}
	return out
}
