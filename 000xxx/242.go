func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)
	for _, b := range []byte(s) {
		m[b-'a']++
	}
	for _, b := range []byte(t) {
		m[b-'a']--
		if m[b-'a'] < 0 {
			return false
		}
	}
	return true
}
