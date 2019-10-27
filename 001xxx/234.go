func balancedString(s string) int {
	l := len(s)
	o, n := l, l>>2
	var cq, cw, ce, cr int
	for i := 0; i < l; i++ {
		switch s[i] {
		case 'Q':
			cq++
		case 'W':
			cw++
		case 'E':
			ce++
		case 'R':
			cr++
		}
	}
	j := 0
	for i := 0; i < l; i++ {
		switch s[i] {
		case 'Q':
			cq--
		case 'W':
			cw--
		case 'E':
			ce--
		case 'R':
			cr--
		}
		for j < l && cq <= n && cw <= n && ce <= n && cr <= n {
			if t := i - j + 1; t < o {
				o = t
			}
			switch s[j] {
			case 'Q':
				cq++
			case 'W':
				cw++
			case 'E':
				ce++
			case 'R':
				cr++
			}
			j++
		}
	}
	return o
}
