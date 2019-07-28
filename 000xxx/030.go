func findSubstring(s string, words []string) []int {
	ls, lws := len(s), len(words)
	if ls == 0 || lws == 0 || ls < lws*len(words[0]) {
		return []int{}
	}
	lw := len(words[0])
	m := map[string]int{}
	for _, w := range words {
		if len(w) != lw {
			return []int{}
		}
		m[w]++
	}
	m1 := map[string]int{}
	c, l, r := 1, 0, 0

	reset := func() {
		if c != 0 {
			for k, v := range m {
				m1[k] = v
			}
			c = 0
		}
	}

	o := make([]int, 0, ls)
	for i := 0; i < lw; i++ {
		l, r = i, i
		reset()
		for ls-l >= lws*lw {
			t := s[r : r+lw]
			if v, ok := m1[t]; !ok {
				l, r = r+lw, r+lw
				reset()
			} else if v == 0 {
				m1[s[l:l+lw]]++
				c, l = c-1, l+lw
			} else {
				m1[t]--
				c, r = c+1, r+lw
				if c == lws {
					m1[s[l:l+lw]]++
					o, c, l = append(o, l), c-1, l+lw
				}
			}
		}
	}
	return o
}
