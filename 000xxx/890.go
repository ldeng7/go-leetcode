func findAndReplacePattern(words []string, pattern string) []string {
	out := []string{}
	cal := func(w string) {
		if len(w) != len(pattern) {
			return
		}
		w2p, p2w := map[byte]byte{}, map[byte]byte{}
		for i := 0; i < len(w); i++ {
			cw, cp := w[i], pattern[i]
			if c, ok := w2p[cw]; ok {
				if cp != c {
					return
				}
			} else {
				if _, ok := p2w[cp]; ok {
					return
				}
				w2p[cw], p2w[cp] = cp, cw
			}
		}
		out = append(out, w)
	}
	for _, w := range words {
		cal(w)
	}
	return out
}
