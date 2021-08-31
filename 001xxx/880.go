func cal(s string) int {
	o := 0
	for i := 0; i < len(s); i++ {
		o = o*10 + int(s[i]-'a')
	}
	return o
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return cal(targetWord) == cal(firstWord)+cal(secondWord)
}
