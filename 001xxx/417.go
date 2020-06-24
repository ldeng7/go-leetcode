func reformat(s string) string {
	l := len(s)
	la, ld := 0, 0
	for i := 0; i < l; i++ {
		if b := s[i]; b >= 'a' {
			la++
		} else {
			ld++
		}
	}
	alphaFirst, d := true, la-ld
	if d < 0 {
		alphaFirst, d = false, -d
	}
	if d > 1 {
		return ""
	}

	bs := make([]byte, l)
	ia, id := 0, 1
	if !alphaFirst {
		ia, id = 1, 0
	}
	for i := 0; i < l; i++ {
		if b := s[i]; b >= 'a' {
			bs[ia] = b
			ia += 2
		} else {
			bs[id] = b
			id += 2
		}
	}
	return string(bs)
}
