func splitString(s string) bool {
	l := len(s)
	var cal func(int, int) bool
	cal = func(i, n int) bool {
		if i == l {
			return true
		}
		for m := 0; i < l; i++ {
			m = m*10 + int(s[i]-'0')
			if m == n-1 && cal(i+1, m) {
				return true
			} else if m >= n {
				return false
			}
		}
		return false
	}

	for i, n := 0, 0; i < l-1; i++ {
		n = n*10 + int(s[i]-'0')
		if cal(i+1, n) {
			return true
		}
	}
	return false
}
