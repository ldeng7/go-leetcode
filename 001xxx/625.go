func min(a, b string) string {
	if a <= b {
		return a
	}
	return b
}

func shift(c, a byte) byte {
	if a == 5 {
		if c < 5 {
			return 0
		}
		return 5
	} else if a&1 == 0 && c&1 == 1 {
		return 11 - c
	}
	return 10 - c
}

func findLexSmallestString(s string, a int, b int) string {
	o, l, b1 := s, len(s), 0
	for {
		bs := []byte(s[b:] + s[:b])
		c := shift(bs[1]-'0', byte(a))
		for i := 1; i < l; i += 2 {
			bs[i] = (bs[i]-'0'+c)%10 + '0'
		}
		if b&1 == 1 {
			c = shift(bs[0]-'0', byte(a))
			for i := 0; i < l; i += 2 {
				bs[i] = (bs[i]-'0'+c)%10 + '0'
			}
		}
		s = string(bs)
		o = min(o, s)
		b1 = (b1 + b) % l
		if b1 == 0 {
			break
		}
	}
	return o
}
