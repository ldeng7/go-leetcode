func parseBoolExpr(expression string) bool {
	i := 0
	var cal func(s string) bool
	cal = func(s string) bool {
		t, b := s[i], 0
		i++
		if t == 't' || t == 'f' {
			return t == 't'
		}
		i++
		for s[i] != ')' {
			if cal(s) {
				b |= 2
			} else {
				b |= 1
			}
			if s[i] == ',' {
				i++
			}
		}
		i++
		if t == '!' {
			return b&1 == 1
		} else if t == '|' {
			return b&2 == 2
		}
		return b&1 == 0
	}
	return cal(expression)
}
