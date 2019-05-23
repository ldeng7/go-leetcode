func lengthOfLastWord(s string) int {
	c := 0
	var i int
	for i = len(s) - 1; i >= 0; i-- {
		if ' ' != s[i] {
			break
		}
	}
	for ; i >= 0; i-- {
		if ' ' == s[i] {
			break
		}
		c++
	}
	return c
}
