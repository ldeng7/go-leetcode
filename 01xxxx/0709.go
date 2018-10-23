func toLowerCase(str string) string {
	bs := make([]byte, len(str))
	for i, c := range []byte(str) {
		if c >= 'A' && c <= 'Z' {
			bs[i] = 'a' + (c - 'A')
		} else {
			bs[i] = c
		}
	}
	return string(bs)
}
