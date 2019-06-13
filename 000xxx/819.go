func mostCommonWord(paragraph string, banned []string) string {
	l, j := len(paragraph), 0
	bs := []byte(paragraph)
	mb := map[string]bool{}
	for _, w := range banned {
		mb[w] = true
	}
	m := map[string]int{}

	for i := 0; i < l; i++ {
		c := bs[i]
		if c >= 'A' && c <= 'Z' {
			bs[i] = c + 32
		}
	}
	for ; j < l; j++ {
		c := bs[j]
		if c >= 'a' && c <= 'z' {
			break
		}
	}
	for i := j; i < l; {
		c := bs[i]
		if c >= 'a' && c <= 'z' {
			i++
		} else {
			w := string(bs[j:i])
			if !mb[w] {
				m[w]++
			}
			for ; i < l; i++ {
				c := bs[i]
				if c >= 'a' && c <= 'z' {
					break
				}
			}
			j = i
		}
	}
	if j != l {
		w := string(bs[j:])
		if !mb[w] {
			m[w]++
		}
	}

	var mw string
	var mc int
	for w, c := range m {
		if c > mc {
			mw, mc = w, c
		}
	}
	return mw
}
