func distinctSubseqII(S string) int {
	o, cs := 0, [26]int{}
	for i := 0; i < len(S); i++ {
		j := S[i] - 'a'
		c := cs[j]
		cs[j] = o + 1
		o = (o<<1 - c + 1000000008) % 1000000007
	}
	return o
}
