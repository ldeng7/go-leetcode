func calculate(s string) int {
	out, sign, l, st := 0, 1, len(s), []int{}
	for i := 0; i < l; i++ {
		c := s[i]
		if c >= '0' {
			n := 0
			for i < l && s[i] >= '0' {
				n = 10*n + int(s[i]-'0')
				i++
			}
			out += sign * n
			i--
		} else if c == '+' {
			sign = 1
		} else if c == '-' {
			sign = -1
		} else if c == '(' {
			st = append(st, out, sign)
			out, sign = 0, 1
		} else if c == ')' {
			j := len(st) - 2
			out *= st[j+1]
			out += st[j]
			st = st[:j]
		}
	}
	return out
}
