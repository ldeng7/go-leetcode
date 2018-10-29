import "strings"

func wordPattern(pattern string, str string) bool {
	p2s := map[byte]string{}
	s2p := map[string]byte{}
	strs := strings.Split(str, " ")
	if len(strs) != len(pattern) {
		return false
	}
	for i, s := range strs {
		p := pattern[i]
		sp, ok := p2s[p]
		if ok {
			if sp != s {
				return false
			}
		} else {
			p2s[p] = s
		}
		ps, ok := s2p[s]
		if ok {
			if ps != p {
				return false
			}
		} else {
			s2p[s] = p
		}
	}
	return true
}
