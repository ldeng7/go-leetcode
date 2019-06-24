func minAddToMakeValid(S string) int {
	c, r := 0, 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			c++
		} else {
			c--
		}
		if c < 0 {
			c, r = 0, r+1
		}
	}
	return c + r
}
