func destCity(paths [][]string) string {
	m := map[string]bool{}
	for _, p := range paths {
		m[p[0]] = true
	}
	for _, p := range paths {
		if s := p[1]; !m[s] {
			return s
		}
	}
	return ""
}
