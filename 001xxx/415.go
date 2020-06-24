func getHappyString(n int, k int) string {
	var o string
	s := make([]byte, 0, n)
	var cal func()
	cal = func() {
		l := len(s)
		if l == n {
			k--
			if k == 0 {
				o = string(s)
			}
			return
		}
		for c := byte('a'); c <= 'c'; c++ {
			if l == 0 || s[l-1] != c {
				s = append(s, c)
				cal()
				s = s[:l]
			}
			if k == 0 {
				return
			}
		}
	}
	cal()
	return o
}
