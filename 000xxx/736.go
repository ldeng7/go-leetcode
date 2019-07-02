import "strconv"

func dupMap(m map[string]int) map[string]int {
	o := map[string]int{}
	for k, v := range m {
		o[k] = v
	}
	return o
}

func evaluate(expression string) int {
	var cal func(string, map[string]int) int
	cal = func(s string, m map[string]int) int {
		if s[0] == '-' || (s[0] >= '0' && s[0] <= '9') {
			i, _ := strconv.Atoi(s)
			return i
		} else if s[0] != '(' {
			return m[s]
		}
		s = s[1 : len(s)-1]
		i := 0

		parse := func() string {
			i0, e, c := i, i+1, 1
			if s[i] == '(' {
				for ; c != 0; e++ {
					if s[e] == '(' {
						c++
					} else if s[e] == ')' {
						c--
					}
				}
			} else {
				for e < len(s) && s[e] != ' ' {
					e++
				}
			}
			i = e + 1
			return s[i0:e]
		}

		switch parse() {
		case "let":
			m1 := dupMap(m)
			for {
				t := parse()
				if i > len(s) {
					return cal(t, m1)
				}
				m1[t] = cal(parse(), m1)
			}
		case "add":
			return cal(parse(), m) + cal(parse(), m)
		case "mult":
			return cal(parse(), m) * cal(parse(), m)
		}
		return 0
	}
	return cal(expression, map[string]int{})
}
