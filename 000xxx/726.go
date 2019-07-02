import (
	"bytes"
	"sort"
	"strconv"
)

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func countOfAtoms(formula string) string {
	i, l := 0, len(formula)
	var cal func() map[string]int
	cal = func() map[string]int {
		m := map[string]int{}
		for i < l {
			c := formula[i]
			if c == '(' {
				i++
				m1 := cal()
				for k, v := range m1 {
					m[k] += v
				}
			} else if c == ')' {
				j := i + 1
				i++
				for i < l && isDigit(formula[i]) {
					i++
				}
				n, _ := strconv.Atoi(formula[j:i])
				for k, _ := range m {
					m[k] *= n
				}
				return m
			} else {
				j := i
				i++
				for i < l && isLower(formula[i]) {
					i++
				}
				k := formula[j:i]
				j = i
				for i < l && isDigit(formula[i]) {
					i++
				}
				n, _ := strconv.Atoi(formula[j:i])
				if n == 0 {
					n = 1
				}
				m[k] += n
			}
		}
		return m
	}

	m := cal()
	ks, i := make([]string, len(m)), 0
	for k, _ := range m {
		ks[i], i = k, i+1
	}
	sort.Strings(ks)
	buf := &bytes.Buffer{}
	for _, k := range ks {
		buf.WriteString(k)
		v := m[k]
		if v != 1 {
			buf.WriteString(strconv.Itoa(v))
		}
	}
	return buf.String()
}
