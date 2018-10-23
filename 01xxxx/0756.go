func cal(cur, top string, m map[string]map[byte]bool) bool {
	if len(cur) == 2 && len(top) == 1 {
		return true
	}
	if len(cur) == len(top)+1 {
		return cal(top, "", m)
	}
	bm, ok := m[cur[len(top):len(top)+2]]
	if ok {
		for b, _ := range bm {
			if cal(cur, top+string(b), m) {
				return true
			}
		}
	}
	return false
}

func pyramidTransition(bottom string, allowed []string) bool {
	m := map[string]map[byte]bool{}
	for _, s := range allowed {
		ss := s[:2]
		if _, ok := m[ss]; !ok {
			m[ss] = map[byte]bool{}
		}
		m[ss][s[2]] = true
	}
	return cal(bottom, "", m)
}
