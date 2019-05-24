func canConstruct(ransomNote string, magazine string) bool {
	m := [26]int{}
	for _, c := range []byte(magazine) {
		m[c-'a']++
	}
	for _, c := range []byte(ransomNote) {
		c -= 'a'
		if m[c] == 0 {
			return false
		}
		m[c]--
	}
	return true
}
