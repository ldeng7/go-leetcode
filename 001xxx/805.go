func numDifferentIntegers(word string) int {
	k, s := 0, map[string]bool{"": true}
	for i := 0; i < len(word); i++ {
		c := word[i]
		if c >= '0' && c <= '9' {
			if k < i && word[k] == '0' {
				k++
			}
		} else {
			s[word[k:i]] = true
			k = i + 1
		}
	}
	s[word[k:]] = true
	return len(s) - 1
}
