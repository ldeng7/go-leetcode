func areSentencesSimilar(words1 []string, words2 []string, pairs [][]string) bool {
	if len(words1) != len(words2) {
		return false
	}
	m := map[string]map[string]bool{}
	for _, p := range pairs {
		if mw, ok := m[p[0]]; !ok {
			m[p[0]] = map[string]bool{p[1]: true}
		} else {
			mw[p[1]] = true
		}
		if mw, ok := m[p[1]]; !ok {
			m[p[1]] = map[string]bool{p[0]: true}
		} else {
			mw[p[0]] = true
		}
	}
	for i := 0; i < len(words1); i++ {
		if words1[i] == words2[i] {
			continue
		}
		if mw, ok := m[words1[i]]; !ok {
			return false
		} else if !mw[words2[i]] {
			return false
		}
	}
	return true
}
