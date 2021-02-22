func countConsistentStrings(allowed string, words []string) int {
	o := len(words)
	t := [26]bool{}
	for i := 0; i < len(allowed); i++ {
		t[allowed[i]-'a'] = true
	}
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			if !t[w[i]-'a'] {
				o--
				break
			}
		}
	}
	return o
}
