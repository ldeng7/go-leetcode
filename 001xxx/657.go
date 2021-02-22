func closeStrings(word1 string, word2 string) bool {
	s, l := 0, len(word1)
	if l != len(word2) {
		return false
	}
	ca, cb := [26]int{}, [26]int{}
	for i := 0; i < l; i++ {
		ca[word1[i]-'a']++
		cb[word2[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if (ca[i] != 0 && cb[i] == 0) || (ca[i] == 0 && cb[i] != 0) {
			return false
		}
		s ^= ca[i] ^ cb[i]
	}
	return s == 0
}
