func makeEqual(words []string) bool {
	l := len(words)
	cs := [26]int{}
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			cs[w[i]-'a']++
		}
	}
	for i := 0; i < 26; i++ {
		if cs[i]%l != 0 {
			return false
		}
	}
	return true
}
