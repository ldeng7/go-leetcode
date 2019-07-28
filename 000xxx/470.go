func rand10() int {
	a, b := rand7(), rand7()
	if a <= 4 || b >= 4 {
		return (a+b)%10 + 1
	}
	return rand10()
}
