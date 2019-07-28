func generateParenthesis(n int) []string {
	o, bs := []string{}, []byte{}
	var cal func(int, int)
	cal = func(a, b int) {
		if a == n && b == n {
			o = append(o, string(bs))
			return
		}
		if b > a || b > n || a > n {
			return
		}
		bs = append(bs, '(')
		cal(a+1, b)
		bs[len(bs)-1] = ')'
		cal(a, b+1)
		bs = bs[:len(bs)-1]
	}
	cal(0, 0)
	return o
}
