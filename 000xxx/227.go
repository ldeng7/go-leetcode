func calculate(s string) int {
	o, l, v, n, op := 0, len(s), 0, 0, byte('+')
	for i := 0; i < l; i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
		if c == '+' || c == '-' || c == '*' || c == '/' || i == l-1 {
			switch op {
			case '+':
				v += n
			case '-':
				v -= n
			case '*':
				v *= n
			case '/':
				v /= n
			}
			if c == '+' || c == '-' || i == l-1 {
				o, v = o+v, 0
			}
			op, n = c, 0
		}
	}
	return o
}
