func minimumSwap(s1 string, s2 string) int {
	l, c1, c2 := len(s1), 0, 0
	for i := 0; i < l; i++ {
		if s1[i] != s2[i] {
			if s1[i] == 'x' {
				c1++
			} else {
				c2++
			}
		}
	}
	if c1&1+c2&1 == 1 {
		return -1
	}
	return (c1+1)/2 + (c2+1)/2
}
