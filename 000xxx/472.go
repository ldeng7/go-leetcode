func findAllConcatenatedWordsInADict(words []string) []string {
	o, s := []string{}, map[string]bool{}
	for _, w := range words {
		s[w] = true
	}
	for _, w := range words {
		l := len(w)
		if l == 0 {
			continue
		}
		t := make([]bool, l+1)
		t[0] = true
		for i := 0; i < l; i++ {
			if !t[i] {
				continue
			}
			for j := i + 1; j <= l; j++ {
				if j-i < l && s[w[i:j]] {
					t[j] = true
				}
			}
			if t[l] {
				o = append(o, w)
				break
			}
		}
	}
	return o
}
