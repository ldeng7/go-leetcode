func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minOperations(s string) int {
	e0, e1, o0, o1 := 0, 0, 0, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if i&1 == 0 {
			if c == '1' {
				e0++
			} else {
				e1++
			}
		} else {
			if c == '1' {
				o0++
			} else {
				o1++
			}
		}
	}
	return min(e1+o0, e0+o1)
}
