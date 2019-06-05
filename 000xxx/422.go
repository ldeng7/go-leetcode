func validWordSquare(words []string) bool {
	n := len(words)
	pl := n
	for i := 0; i < n; i++ {
		s := words[i]
		if len(s) > pl {
			return false
		}
		if len(s) < len(words) && len(words[len(s)]) > i {
			return false
		}
		for j := 0; j < len(s); j++ {
			if len(words[j]) <= i || s[j] != words[j][i] {
				return false
			}
		}
		pl = len(s)
	}
	return true
}
