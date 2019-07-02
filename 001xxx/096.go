import "sort"

func merge(a, b map[string]bool) {
	for s, _ := range b {
		a[s] = true
	}
}

func mul(a, b map[string]bool) map[string]bool {
	if 0 == len(a) {
		return b
	}
	c := map[string]bool{}
	for sa, _ := range a {
		for sb, _ := range b {
			c[sa+sb] = true
		}
	}
	return c
}

func braceExpansionII(expression string) []string {
	l := len(expression)
	var cal func(int) (map[string]bool, int)
	cal = func(i int) (map[string]bool, int) {
		s, p := map[string]bool{}, map[string]bool{}
		for i < l {
			c := expression[i]
			switch c {
			case ',':
				merge(s, p)
				p = map[string]bool{}
				i++
			case '{':
				s1, j := cal(i + 1)
				p, i = mul(p, s1), j
			case '}':
				merge(s, p)
				return s, i + 1
			default:
				bs := []byte{}
				for i < l && expression[i] >= 'a' && expression[i] <= 'z' {
					bs, i = append(bs, expression[i]), i+1
				}
				p = mul(p, map[string]bool{string(bs): true})
			}
		}
		merge(s, p)
		return s, i
	}

	s, _ := cal(0)
	o, i := make([]string, len(s)), 0
	for str, _ := range s {
		o[i], i = str, i+1
	}
	sort.Strings(o)
	return o
}
