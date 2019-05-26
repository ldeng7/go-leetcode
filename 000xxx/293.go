func generatePossibleNextMoves(s string) []string {
	out := []string{}
	for i := 0; i < len(s)-1; i++ {
		if '+' == s[i] && '+' == s[i+1] {
			bs := []byte(s)
			bs[i], bs[i+1] = '-', '-'
			out = append(out, string(bs))
		}
	}
	return out
}
