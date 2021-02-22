func reformatNumber(number string) string {
	l := len(number)
	ar := make([]byte, 0, l)
	for i := 0; i < l; i++ {
		if c := number[i]; c >= '0' && c <= '9' {
			ar = append(ar, c)
		}
	}
	l = len(ar)
	o := make([]byte, 0, l/3*4)
	for i := 0; i < l; i++ {
		if i != 0 && i%3 == 0 {
			if i < l-1 {
				o = append(o, '-')
			} else {
				j := len(o) - 1
				c := o[j]
				o[j] = '-'
				o = append(o, c)
			}
		}
		o = append(o, ar[i])
	}
	return string(o)
}
