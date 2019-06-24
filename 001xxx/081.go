func smallestSubsequence(text string) string {
	l := len(text)
	pre, post, st := 0, make([]int, l), []byte{}
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			post[i] |= 1 << (text[j] - 'a')
		}
	}
	for i := 0; i < l; i++ {
		c := text[i]
		if pre&(1<<(c-'a')) != 0 {
			continue
		}
		es := len(st) - 1
		for 0 != len(st) && c < st[es] && post[i]&(1<<(st[es]-'a')) != 0 {
			pre ^= 1 << (st[es] - 'a')
			st, es = st[:es], es-1
		}
		pre |= 1 << (c - 'a')
		st = append(st, c)
	}
	return string(st)
}
