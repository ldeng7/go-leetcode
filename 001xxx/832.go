func checkIfPangram(sentence string) bool {
	m := [26]bool{}
	for i := 0; i < len(sentence); i++ {
		m[sentence[i]-'a'] = true
	}
	for i := 0; i < 26; i++ {
		if !m[i] {
			return false
		}
	}
	return true
}
