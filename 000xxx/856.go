func scoreOfParentheses(S string) int {
	o, c := 0, uint(0)
	for i, b := range []byte(S) {
		if b == '(' {
			c++
		} else {
			c--
			if b == ')' && S[i-1] == '(' {
				o += 1 << c
			}
		}
	}
	return o
}
