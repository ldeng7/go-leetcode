func modifyString(s string) string {
	bs := []byte(s)
	l := len(bs)
	if l == 1 {
		if bs[0] == '?' {
			return "a"
		}
		return s
	}

	if bs[0] == '?' {
		if bs[1] != 'a' {
			bs[0] = 'a'
		} else {
			bs[0] = 'b'
		}
	}
	for i := 1; i < l-1; i++ {
		if bs[i] == '?' {
			p, s := bs[i-1], bs[i+1]
			if p != 'a' {
				if s != 'a' {
					bs[i] = 'a'
				} else if p != 'b' {
					bs[i] = 'b'
				} else {
					bs[i] = 'c'
				}
			} else {
				if s != 'b' {
					bs[i] = 'b'
				} else {
					bs[i] = 'c'
				}
			}
		}
	}
	if bs[l-1] == '?' {
		if bs[l-2] != 'a' {
			bs[l-1] = 'a'
		} else {
			bs[l-1] = 'b'
		}
	}
	return string(bs)
}
