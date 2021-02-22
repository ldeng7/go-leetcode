func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxDepth(s string) int {
	o, c, l := 0, 0, len(s)
	for i := 0; i < l; i++ {
		if b := s[i]; b == '(' {
			c++
		} else if b == ')' {
			c--
		}
		o = max(o, c)
	}
	return o
}
