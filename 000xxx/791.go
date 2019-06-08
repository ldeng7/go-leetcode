func customSortString(S string, T string) string {
	ls, lt := len(S), len(T)
	cs := [26]int{}
	bs := make([]byte, lt)
	j := 0
	for i := 0; i < lt; i++ {
		cs[T[i]-'a']++
	}
	var c int
	for i := 0; i < ls; i++ {
		k := S[i] - 'a'
		cs[k], c = 0, cs[k]
		for ; c > 0; c-- {
			bs[j], j = S[i], j+1
		}
	}
	for i := 0; i < lt; i++ {
		k := T[i] - 'a'
		cs[k], c = 0, cs[k]
		for ; c > 0; c-- {
			bs[j], j = T[i], j+1
		}
	}
	return string(bs)
}
