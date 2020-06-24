func numSteps(s string) int {
	a, b := 0, 0
	for i := len(s) - 1; i > 0; i-- {
		if int(s[i]-'0')+b == 1 {
			a, b = a+2, 1
		} else {
			a += 1
		}
	}
	return a + b
}
