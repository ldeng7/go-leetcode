func balancedStringSplit(s string) int {
	c0, c1 := 0, 0
	o, l := 0, len(s)
	for i := 0; i < l; i++ {
		if s[i] == 'L' {
			c0++
		} else {
			c1++
		}
		if c0 == c1 {
			o, c0, c1 = o+1, 0, 0
		}
	}
	return o
}
