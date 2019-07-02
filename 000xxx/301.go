func removeInvalidParentheses(s string) []string {
	o := []string{}
	var cal func([]byte, int, int, byte, byte)
	cal = func(bs []byte, i0, j0 int, p0, p1 byte) {
		c, i := 0, i0
		for ; i < len(bs) && c >= 0; i++ {
			if bs[i] == p0 {
				c++
			} else if bs[i] == p1 {
				c--
			}
		}
		if c < 0 {
			i--
			for j := j0; j <= i; j++ {
				if bs[j] == p1 && (j == j0 || bs[j] != bs[j-1]) {
					bs1 := make([]byte, len(bs)-1)
					copy(bs1, bs[:j])
					copy(bs1[j:], bs[j+1:])
					cal(bs1, i, j, p0, p1)
				}
			}
		} else {
			l := len(bs)
			for i = 0; i < l>>1; i++ {
				bs[i], bs[l-i-1] = bs[l-i-1], bs[i]
			}
			if p0 == '(' {
				cal(bs, 0, 0, ')', '(')
			} else {
				o = append(o, string(bs))
			}
		}
	}
	cal([]byte(s), 0, 0, '(', ')')
	return o
}
