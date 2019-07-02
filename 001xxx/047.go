func removeDuplicates(S string) string {
	st, l := []byte{}, len(S)
	for i := 0; i < l; i++ {
		c, e := S[i], len(st)-1
		if e >= 0 && st[e] == c {
			st = st[:e]
		} else {
			st = append(st, c)
		}
	}
	return string(st)
}
