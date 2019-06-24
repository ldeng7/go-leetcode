func camelMatch(queries []string, pattern string) []bool {
	lp := len(pattern)
	o := make([]bool, len(queries))
	cal := func(q string) bool {
		j := 0
		for k := 0; k < len(q); k++ {
			c := q[k]
			if j < lp && c == pattern[j] {
				j++
			} else if c >= 'A' && c <= 'Z' {
				return false
			}
		}
		return j == lp
	}
	for i, q := range queries {
		o[i] = cal(q)
	}
	return o
}
