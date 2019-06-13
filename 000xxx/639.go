func numDecodings(s string) int {
	e0, e1, e2, f0 := 1, 0, 0, 0
	for _, c := range []byte(s) {
		if c == '*' {
			f0 = 9*e0 + 9*e1 + 6*e2
			e1, e2 = e0, e0
		} else {
			f0 = e1
			if c > '0' {
				f0 += e0
			}
			if c <= '6' {
				f0 += e2
			}
			e1, e2 = 0, 0
			if c == '1' {
				e1 = e0
			} else if c == '2' {
				e2 = e0
			}
		}
		e0 = f0 % 1000000007
	}
	return e0
}
