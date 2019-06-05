func areSentencesSimilarTwo(words1 []string, words2 []string, pairs [][]string) bool {
	if len(words1) != len(words2) {
		return false
	}
	m := map[string]string{}
	var cal func(string) string
	cal = func(s string) string {
		if _, ok := m[s]; !ok {
			m[s] = s
		}
		if s == m[s] {
			return s
		}
		return cal(m[s])
	}
	for _, p := range pairs {
		r1, r2 := cal(p[0]), cal(p[1])
		if r1 != r2 {
			m[r1] = r2
		}
	}
	for i := 0; i < len(words1); i++ {
		if cal(words1[i]) != cal(words2[i]) {
			return false
		}
	}
	return true
}
