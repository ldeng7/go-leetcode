func freqAlphabets(s string) string {
	l := len(s)
	o := make([]byte, 0, l)
	for i := 0; i < l; i++ {
		c := s[i] - '0'
		if i < l-2 && s[i+2] == '#' {
			c = c*10 + s[i+1] - '0'
			i += 2
		}
		o = append(o, c+'`')
	}
	return string(o)
}
