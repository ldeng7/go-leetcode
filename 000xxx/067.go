func addBinary(a string, b string) string {
	s := ""
	la, lb := len(a)-1, len(b)-1
	var c byte = 0
	for la >= 0 || lb >= 0 {
		var da, db byte = 0, 0
		if la >= 0 {
			da = a[la] - '0'
			la--
		}
		if lb >= 0 {
			db = b[lb] - '0'
			lb--
		}
		d := da + db + c
		if d&2 == 2 {
			d &= 1
			c = 1
		} else {
			c = 0
		}
		s = string(d+'0') + s
	}
	if 1 == c {
		s = "1" + s
	}
	return s
}
