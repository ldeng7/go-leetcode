func reverseWords(str []byte) {
	var cal func(l, r int)
	cal = func(l, r int) {
		for l < r {
			str[l], str[r] = str[r], str[l]
			l, r = l+1, r-1
		}
	}
	l := 0
	for i := 0; i <= len(str); i++ {
		if i == len(str) || str[i] == ' ' {
			cal(l, i-1)
			l = i + 1
		}
	}
	cal(0, len(str)-1)
}
