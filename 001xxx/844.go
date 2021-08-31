func replaceDigits(s string) string {
	bs := []byte(s)
	for i := 1; i < len(s); i += 2 {
		bs[i] = bs[i-1] + (bs[i] - '0')
	}
	return string(bs)
}
