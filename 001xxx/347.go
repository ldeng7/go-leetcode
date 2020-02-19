func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func minSteps(s string, t string) int {
	ls, lt := len(s), len(t)
	ars, art := [26]int{}, [26]int{}
	for i := 0; i < ls; i++ {
		ars[s[i]-'a']++
	}
	for i := 0; i < lt; i++ {
		art[t[i]-'a']++
	}
	o := 0
	for i := 0; i < 26; i++ {
		o += abs(ars[i] - art[i])
	}
	return o >> 1
}
