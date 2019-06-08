func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxProduct(words []string) int {
	out := 0
	m := map[int]int{}
	for _, w := range words {
		mask := 0
		for _, c := range []byte(w) {
			mask |= 1 << (c - 'a')
		}
		m[mask] = max(m[mask], len(w))
		for mask1, l := range m {
			if mask&mask1 == 0 {
				out = max(out, len(w)*l)
			}
		}
	}
	return out
}
