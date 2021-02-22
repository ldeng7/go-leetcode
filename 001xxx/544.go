func toLower(c byte) byte {
	if c >= 'a' {
		return c
	}
	return 'a' + c - 'A'
}

func makeGood(s string) string {
	l := len(s)
	bs := make([]byte, 0, l)
	for i := 0; i < l; i++ {
		c := s[i]
		if 0 != len(bs) && toLower(bs[len(bs)-1]) == toLower(c) && bs[len(bs)-1] != c {
			bs = bs[:len(bs)-1]
		} else {
			bs = append(bs, c)
		}
	}
	return string(bs)
}
