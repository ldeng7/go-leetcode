func rand10() int {
	a := (rand7()-1)*7 + rand7()
	if a <= 40 {
		return a%10 + 1
	}
	return rand10()
}
