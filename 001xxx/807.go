import "bytes"

func evaluate(s string, knowledge [][]string) string {
	l := len(s)
	o := bytes.NewBuffer(make([]byte, 0, l*2))
	m := map[string]string{}
	for _, k := range knowledge {
		m[k[0]] = k[1]
	}
	for i := 0; i < l; i++ {
		if c := s[i]; c != '(' {
			o.WriteByte(c)
		} else {
			j := i + 1
			for ; s[j] != ')'; j++ {
			}
			if t, ok := m[s[i+1:j]]; ok {
				o.WriteString(t)
			} else {
				o.WriteByte('?')
			}
			i = j
		}
	}
	return o.String()
}
