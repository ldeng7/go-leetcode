func myAtoi(str string) int {
	ds := []byte{}
	m := false
	started := false
	for _, c := range []byte(str) {
		if c == ' ' {
			if started {
				break
			}
			continue
		} else if c == '+' || c == '-' {
			if !started {
				started = true
				m = c == '-'
				continue
			}
			break
		} else if c < '0' || c > '9' {
			break
		}
		started = true
		ds = append(ds, c)
	}

	dst := []byte{}
	started = false
	for _, c := range ds {
		if c == '0' && !started {
			continue
		}
		started = true
		dst = append(dst, c)
	}
	ds = dst
	if len(ds) > 11 {
		if !m {
			return 2147483647
		} else {
			return -2147483648
		}
	}

	var o, b uint64 = 0, 1
	for i := len(ds) - 1; i >= 0; i-- {
		o += uint64(ds[i]-'0') * b
		b *= 10
		if !m {
			if o >= 2147483648 {
				return 2147483647
			}
		} else {
			if o > 2147483648 {
				return -2147483648
			}
		}
	}
	i := int(o)
	if m {
		i = -i
	}
	return i
}