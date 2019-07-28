func titleToNumber(s string) int {
	o, b := 0, 1
	for i := len(s) - 1; i >= 0; i-- {
		o += int(s[i]-'A'+1) * b
		b *= 26
	}
	return o
}
