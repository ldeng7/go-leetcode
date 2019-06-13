func numberOfLines(widths []int, S string) []int {
	l, w := 1, 0
	for _, c := range []byte(S) {
		w1 := widths[c-'a']
		if w+w1 > 100 {
			l, w = l+1, w1
		} else {
			w += w1
		}
	}
	return []int{l, w}
}
