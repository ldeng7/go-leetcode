func generateParenthesis(n int) []string {
	strs := []string{}
	var gen func(int, int, int, string)
	gen = func(a, b, n int, s string) {
		if a == n && b == n {
			strs = append(strs, s)
			return
		}
		if b > a || b > n || a > n {
			return
		}
		gen(a+1, b, n, s+"(")
		gen(a, b+1, n, s+")")
	}
	gen(0, 0, n, "")
	return strs
}
