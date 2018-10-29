func findSubstring(s string, words []string) []int {
	if 0 == len(s) || 0 == len(words) {
		return []int{}
	}
	out := []int{}
	l, l0 := len(words), len(words[0])
	m := map[string]int{}
	for _, w := range words {
		m[w] = m[w] + 1
	}
	for i := 0; i <= len(s)-l*l0; i++ {
		mm := map[string]int{}
		j := 0
		for ; j < l; j++ {
			t := s[i+j*l0 : i+(j+1)*l0]
			if _, ok := m[t]; !ok {
				break
			}
			mm[t] = mm[t] + 1
			if mm[t] > m[t] {
				break
			}
		}
		if j == l {
			out = append(out, i)
		}
	}
	return out
}
