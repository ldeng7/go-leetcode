func calculate(s string) int {
	l, n, o, out := len(s), 0, 0, 0
	var op byte = '+'
	for i := 0; i < l; i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		} else if c == '(' {
			j, cnt := i, 0
			for ; i < l; i++ {
				if s[i] == '(' {
					cnt++
				} else if s[i] == ')' {
					cnt--
				}
				if cnt == 0 {
					break
				}
			}
			n = calculate(s[j+1 : i])
		}
		b := c == '+' || c == '-' || i == l-1
		if b || c == '*' || c == '/' {
			switch op {
			case '+':
				o += n
			case '-':
				o -= n
			case '*':
				o *= n
			case '/':
				o /= n
			}
			if b {
				out, o = out+o, 0
			}
			op, n = c, 0
		}
	}
	return out
}
