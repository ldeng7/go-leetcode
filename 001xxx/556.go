func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}
	bs := make([]byte, 0, 13)
	for c := 0; n != 0; n, c = n/10, c+1 {
		if c == 3 {
			bs = append(bs, '.')
			c = 0
		}
		bs = append(bs, '0'+byte(n%10))
	}
	l := len(bs)
	bs1 := make([]byte, l)
	for i := 0; i < l; i++ {
		bs1[i] = bs[l-i-1]
	}
	return string(bs1)
}
