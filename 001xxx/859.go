func sortSentence(s string) string {
	ar := [9][2]int{}
	l, c := 0, 0
	for i := 1; i < len(s); i++ {
		b := s[i]
		if b < '0' || b > '9' {
			continue
		}
		ar[b-'0'-1] = [2]int{l, i}
		l, c, i = i+2, c+1, i+1
	}

	bs := make([]byte, len(s))
	j := 0
	for i := 0; i < c; i++ {
		l, r := ar[i][0], ar[i][1]
		copy(bs[j:], s[l:r])
		j += r - l
		bs[j] = ' '
		j++
	}
	return string(bs[:j-1])
}
