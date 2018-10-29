func numJewelsInStones(j string, s string) int {
	m := map[byte]bool{}
	for _, b := range []byte(j) {
		m[b] = true
	}
	out := 0
	for _, b := range []byte(s) {
		if m[b] {
			out++
		}
	}
	return out
}
