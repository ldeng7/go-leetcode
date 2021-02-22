func stringShift(s string, shift [][]int) string {
	c := 0
	for _, s := range shift {
		if s[0] == 1 {
			c += s[1]
		} else {
			c -= s[1]
		}
	}
	l := len(s)
	c = (c%l + l) % l
	return s[l-c:] + s[:l-c]
}
