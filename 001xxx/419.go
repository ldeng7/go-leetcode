func minNumberOfFrogs(croakOfFrogs string) int {
	var a, b, c, d, e int
	l := len(croakOfFrogs)
	for i := 0; i < l; i++ {
		switch croakOfFrogs[i] {
		case 'c':
			a++
			if e != 0 {
				e--
			}
		case 'r':
			if a <= 0 {
				return -1
			}
			a, b = a-1, b+1
		case 'o':
			if b <= 0 {
				return -1
			}
			b, c = b-1, c+1
		case 'a':
			if c <= 0 {
				return -1
			}
			c, d = c-1, d+1
		case 'k':
			if d <= 0 {
				return -1
			}
			d, e = d-1, e+1
		}
	}
	if a != 0 || b != 0 || c != 0 || d != 0 {
		return -1
	}
	return e
}
