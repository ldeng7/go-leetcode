func div(a, b int) (int, int) {
	return a / b, a % b
}

func reorderSpaces(text string) string {
	l, ns, b := len(text), 0, -1
	ar := make([][2]int, 0, 16)
	for i := 0; i < l; i++ {
		if c := text[i]; c == ' ' {
			ns++
			if b != -1 {
				ar = append(ar, [2]int{b, i})
				b = -1
			}
		} else if b == -1 {
			b = i
		}
	}
	if b != -1 {
		ar = append(ar, [2]int{b, l})
	}

	bs := make([]byte, l)
	var i, nt int
	if len(ar) != 1 {
		ns, nt = div(ns, len(ar)-1)
		for j, w := range ar {
			copy(bs[i:], []byte(text[w[0]:w[1]]))
			i += w[1] - w[0]
			if j != len(ar)-1 {
				for k := 0; k < ns; k++ {
					bs[i+k] = ' '
				}
				i += ns
			}
		}
	} else {
		b, e := ar[0][0], ar[0][1]
		copy(bs, []byte(text[b:e]))
		i, nt = e-b, ns
	}
	for k := 0; k < nt; k++ {
		bs[i+k] = ' '
	}
	return string(bs)
}
