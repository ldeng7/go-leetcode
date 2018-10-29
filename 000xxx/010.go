func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	if 0 == lp {
		return 0 == ls
	}
	if lp > 1 && p[1] == '*' {
		return isMatch(s, p[2:]) || (ls > 0 && (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p))
	} else {
		return ls > 0 && (s[0] == p[0] || p[0] == '.') && isMatch(s[1:], p[1:])
	}
}
