func wordsTyping(sentence []string, rows int, cols int) int {
	out, idx, l := 0, 0, 0
	for _, s := range sentence {
		l += len(s) + 1
	}
	for i := 0; i < rows; i++ {
		cs := cols
		for cs > 0 {
			if cs >= len(sentence[idx]) {
				cs -= len(sentence[idx])
				if cs > 0 {
					cs--
				}
				idx++
				if idx >= len(sentence) {
					out += cs/l + 1
					cs %= l
					idx = 0
				}
			} else {
				break
			}
		}
	}
	return out
}
