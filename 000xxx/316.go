func removeDuplicateLetters(s string) string {
	m, b := [26]int{}, [26]bool{}
	bs, out := []byte(s), []byte{0}
	for _, c := range bs {
		m[c-'a']++
	}
	for _, c := range bs {
		i := c - 'a'
		m[i]--
		if b[i] {
			continue
		}
		for cl := out[len(out)-1]; c < cl && m[cl-'a'] != 0; cl = out[len(out)-1] {
			i := cl - 'a'
			b[i], out = false, out[:len(out)-1]
		}
		b[i], out = true, append(out, c)
	}
	return string(out[1:])
}
