func canWin(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == '+' && s[i-1] == '+' && !canWin(s[:i-1]+"--"+s[i+1:]) {
			return true
		}
	}
	return false
}
