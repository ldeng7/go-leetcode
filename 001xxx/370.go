func sortString(s string) string {
	l := len(s)
	bs := make([]byte, l)
	cnts := [26]int{}
	for i := 0; i < l; i++ {
		cnts[s[i]-'a']++
	}
	for i := 0; i < l; {
		for j := 0; j < 26; j++ {
			if cnts[j] > 0 {
				bs[i] = 'a' + byte(j)
				i++
				cnts[j]--
			}
		}
		for j := 25; j >= 0; j-- {
			if cnts[j] > 0 {
				bs[i] = 'a' + byte(j)
				i++
				cnts[j]--
			}
		}
	}
	return string(bs)
}
