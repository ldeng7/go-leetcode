func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}
	var m [26]int
	for _, c := range []byte(s1) {
		m[c-'a']++
	}
	for _, c := range []byte(s2) {
		m[c-'a']--
		if m[c-'a'] < 0 {
			return false
		}
	}
	l := len(s1)
	for i := 1; i < l; i++ {
		if (isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:])) ||
			(isScramble(s1[:i], s2[l-i:]) && isScramble(s1[i:], s2[:l-i])) {
			return true
		}
	}
	return false
}
