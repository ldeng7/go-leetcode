func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func countBinarySubstrings(s string) int {
	if 0 == len(s) {
		return 0
	}
	sum, c0, c, b0 := 0, 0, 0, s[0]
	for _, b := range []byte(s) {
		if b != b0 {
			sum += min(c, c0)
			c0, c = c, 1
		} else {
			c++
		}
		b0 = b
	}
	return sum + min(c, c0)
}
