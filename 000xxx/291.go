func wordPatternMatch(pattern string, str string) bool {
	m := map[byte]string{}
	s := map[string]bool{}

	var cal func(string, string) bool
	cal = func(pattern string, str string) bool {
		if 0 == len(pattern) {
			return 0 == len(str)
		}
		if t, ok := m[pattern[0]]; ok {
			if len(t) > len(str) || str[:len(t)] != t {
				return false
			}
			if cal(pattern[1:], str[len(t):]) {
				return true
			}
		} else {
			for i := 1; i <= len(str); i++ {
				if s[str[:i]] {
					continue
				}
				m[pattern[0]] = str[:i]
				s[str[:i]] = true
				if cal(pattern[1:], str[i:]) {
					return true
				}
				delete(m, pattern[0])
				s[str[:i]] = false
			}
		}
		return false
	}
	return cal(pattern, str)
}
