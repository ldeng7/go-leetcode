func numMatchingSubseq(S string, words []string) int {
	o, l := 0, len(S)
	s1, s2 := map[string]bool{}, map[string]bool{}
	for _, w := range words {
		if s1[w] {
			o++
			continue
		} else if s2[w] {
			continue
		}
		i, j, n := 0, 0, len(w)
		for i < l && j < n {
			if w[j] == S[i] {
				j++
			}
			i++
		}
		if j == n {
			o++
			s1[w] = true
		} else {
			s2[w] = true
		}
	}
	return o
}
