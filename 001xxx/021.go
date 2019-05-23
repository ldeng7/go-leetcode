func removeOuterParentheses(S string) string {
	bs := []byte{}
	d, ok := 0, false
	for _, c := range []byte(S) {
		if c == '(' {
			d++
			ok = d > 1
		} else {
			ok = d > 1
			d--
		}
		if ok {
			bs = append(bs, c)
		}
	}
	return string(bs)
}
