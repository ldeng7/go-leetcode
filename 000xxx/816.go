func cal(s string) []string {
	l := len(s)
	if l == 0 || (l > 1 && s[0] == '0' && s[l-1] == '0') {
		return []string{}
	}
	if l > 1 && s[0] == '0' {
		return []string{"0." + s[1:]}
	}
	if s[l-1] == '0' {
		return []string{s}
	}
	out := make([]string, l)
	out[0] = s
	for i := 1; i < l; i++ {
		out[i] = s[:i] + "." + s[i:]
	}
	return out
}

func ambiguousCoordinates(S string) []string {
	out := []string{}
	l := len(S) - 2
	for i := 1; i < l; i++ {
		ssa, ssb := cal(S[1:1+i]), cal(S[i+1:l+1])
		for _, sa := range ssa {
			for _, sb := range ssb {
				out = append(out, "("+sa+", "+sb+")")
			}
		}
	}
	return out
}
