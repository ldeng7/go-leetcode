func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func calculateTime(keyboard string, word string) int {
	m := [26]int{}
	for i := 0; i < 26; i++ {
		m[keyboard[i]-'a'] = i
	}
	o, p := 0, 0
	for i := 0; i < len(word); i++ {
		j := m[word[i]-'a']
		o += abs(p - j)
		p = j
	}
	return o
}
