func addBinary(a string, b string) string {
	bs := []byte{}
	la, lb := len(a)-1, len(b)-1
	var c byte = 0
	for la >= 0 || lb >= 0 {
		var da, db byte = 0, 0
		if la >= 0 {
			da, la = a[la]-'0', la-1
		}
		if lb >= 0 {
			db, lb = b[lb]-'0', lb-1
		}
		d := da + db + c
		if d&2 == 2 {
			d, c = d&1, 1
		} else {
			c = 0
		}
		bs = append(bs, d+'0')
	}
	if 1 == c {
		bs = append(bs, '1')
	}
	l := len(bs)
	for i := 0; i < l>>1; i++ {
		bs[i], bs[l-i-1] = bs[l-i-1], bs[i]
	}
	return string(bs)
}
