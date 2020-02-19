func removePalindromeSub(s string) int {
	l := len(s)
	if 0 == l {
		return 0
	}
	o, i := 2, 0
	for ; i < l>>1; i++ {
		if s[i] != s[l-i-1] {
			break
		}
	}
	if i >= l>>1 {
		o--
	}
	return o
}
