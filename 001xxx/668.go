func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxRepeating(sequence string, word string) int {
	o, ls, lw := 0, len(sequence), len(word)
	if lw > ls {
		return 0
	}
	for i := 0; i <= ls-lw; i++ {
		if sequence[i:i+lw] != word {
			continue
		}
		m := 1
		for j := i + lw; j <= ls-lw; j += lw {
			if sequence[j:j+lw] == word {
				m++
			} else {
				break
			}
		}
		o = max(o, m)
	}
	return o
}
