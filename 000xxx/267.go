import "bytes"

func generatePalindromes(s string) []string {
	m := map[byte]int{}
	t, mid := []byte{}, []byte{}
	for _, b := range []byte(s) {
		m[b]++
	}
	for b, c := range m {
		if c&1 == 1 {
			mid = append(mid, b)
		}
		m[b] >>= 1
		for i := 0; i < m[b]; i++ {
			t = append(t, b)
		}
		if len(mid) > 1 {
			return []string{}
		}
	}

	out := []string{}
	var cal func([]byte)
	cal = func(s []byte) {
		if len(s) >= len(t) {
			buf := &bytes.Buffer{}
			buf.Write(s)
			buf.Write(mid)
			for i := len(s) - 1; i >= 0; i-- {
				buf.WriteByte(s[i])
			}
			out = append(out, buf.String())
			return
		}
		for b, c := range m {
			if c > 0 {
				m[b]--
				cal(append(s, b))
				m[b]++
			}
		}
	}
	cal([]byte{})
	return out
}
