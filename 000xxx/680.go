func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	check := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l, r = l+1, r-1
		}
		return true
	}
	for l < r {
		if s[l] != s[r] {
			return check(l, r-1) || check(l+1, r)
		}
		l, r = l+1, r-1
	}
	return true
}
