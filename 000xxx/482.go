func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}

func licenseKeyFormatting(S string, K int) string {
	out, s := []byte{}, []byte(S)
	i, c := len(s)-1, 0
	for ; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}
		out = append(out, toUpper(s[i]))
		c++
		if c == K {
			out, c = append(out, '-'), 0
		}
	}
	if len(out) != 0 && out[len(out)-1] == '-' {
		out = out[:len(out)-1]
	}
	for i := 0; i < len(out)>>1; i++ {
		out[i], out[len(out)-i-1] = out[len(out)-i-1], out[i]
	}
	return string(out)
}
