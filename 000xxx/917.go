func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func reverseOnlyLetters(S string) string {
	bs := []byte(S)
	i0, i1 := 0, len(bs)-1
	for i0 < i1 {
		for i0 != i1 && !isLetter(bs[i0]) {
			i0++
		}
		for i1 != i0 && !isLetter(bs[i1]) {
			i1--
		}
		bs[i0], bs[i1] = bs[i1], bs[i0]
		i0++
		i1--
	}
	return string(bs)
}
