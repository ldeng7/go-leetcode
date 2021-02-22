func decrypt(code []int, k int) []int {
	l := len(code)
	o := make([]int, l)
	if k > 0 {
		s, b, e := 0, 1, k
		for i := b; i <= e; i++ {
			s += code[i]
		}
		o[0] = s
		for i := 1; i < l; i++ {
			s -= code[b]
			b, e = b+1, e+1
			if b == l {
				b = 0
			}
			if e == l {
				e = 0
			}
			s += code[e]
			o[i] = s
		}
	} else if k < 0 {
		k = -k
		s, b, e := 0, l-2, l-k-1
		for i := b; i >= e; i-- {
			s += code[i]
		}
		o[l-1] = s
		for i := l - 2; i >= 0; i-- {
			s -= code[b]
			b, e = b-1, e-1
			if b == -1 {
				b = l - 1
			}
			if e == -1 {
				e = l - 1
			}
			s += code[e]
			o[i] = s
		}
	}
	return o
}
