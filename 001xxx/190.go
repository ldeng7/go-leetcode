func reverseParentheses(s string) string {
	i, l := 0, len(s)
	var cal func() []byte
	cal = func() []byte {
		o := make([]byte, 0, len(s))
		for i < l && s[i] != ')' {
			if s[i] == '(' {
				i++
				t := cal()
				i++
				lt := len(t)
				for j := 0; j < lt>>1; j++ {
					t[j], t[lt-j-1] = t[lt-j-1], t[j]
				}
				o = append(o, t...)
			} else {
				o = append(o, s[i])
				i++
			}
		}
		return o
	}
	return string(cal())
}
