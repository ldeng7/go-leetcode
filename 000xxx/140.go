func wordBreak(s string, wordDict []string) []string {
	m := map[string][]string{}
	var cal func(string) []string
	cal = func(s string) []string {
		if ar, ok := m[s]; ok {
			return ar
		} else if 0 == len(s) {
			return []string{""}
		}
		o := []string{}
		for _, w := range wordDict {
			l := len(w)
			if len(s) < l {
				l = len(s)
			}
			if s[:l] != w {
				continue
			}
			ar := cal(s[len(w):])
			for _, t := range ar {
				s1 := w
				if 0 != len(t) {
					s1 += " "
					s1 += t
				}
				o = append(o, s1)
			}
		}
		m[s] = o
		return o
	}
	return cal(s)
}
