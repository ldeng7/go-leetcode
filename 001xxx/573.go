func numWays(s string) int {
	l, c := len(s), 0
	for i := 0; i < l; i++ {
		if s[i] == '1' {
			c++
		}
	}
	if c == 0 {
		return ((l - 1) * (l - 2) / 2) % 1000000007
	} else if c%3 != 0 {
		return 0
	}
	c /= 3
	k, c1, c2 := 0, 1, 1
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			k++
			if k > 2*c {
				break
			}
		} else {
			if k == c {
				c1++
			} else if k == c*2 {
				c2++
			}
		}
	}
	return c1 * c2 % 1000000007
}
