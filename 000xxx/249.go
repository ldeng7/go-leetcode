func groupStrings(strings []string) [][]string {
	m := map[string][]string{}
	for _, s := range strings {
		k := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			k[i] = (26 + s[i] - s[0]) % 26
		}
		ks := string(k)
		if ar, ok := m[ks]; ok {
			m[ks] = append(ar, s)
		} else {
			m[ks] = []string{s}
		}
	}
	out := make([][]string, len(m))
	i := 0
	for _, ar := range m {
		out[i] = ar
		i++
	}
	return out
}
