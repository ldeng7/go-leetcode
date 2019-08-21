func countCharacters(words []string, chars string) int {
	o, t := 0, [26]int{}
	for i := 0; i < len(chars); i++ {
		t[chars[i]-'a']++
	}
	for _, w := range words {
		t1, l := t, len(w)
		for i := 0; i < len(w); i++ {
			j := w[i] - 'a'
			t1[j]--
			if t1[j] == -1 {
				l = 0
				break
			}
		}
		o += l
	}
	return o
}
