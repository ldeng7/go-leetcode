var l1 = map[byte]bool{'q': true, 'w': true, 'e': true, 'r': true, 't': true, 'y': true, 'u': true, 'i': true, 'o': true, 'p': true}
var l2 = map[byte]bool{'a': true, 's': true, 'd': true, 'f': true, 'g': true, 'h': true, 'j': true, 'k': true, 'l': true}
var l3 = map[byte]bool{'z': true, 'x': true, 'c': true, 'v': true, 'b': true, 'n': true, 'm': true}

func lower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return 'a' + b - 'A'
	}
	return b
}

func findWords(words []string) []string {
	out := []string{}
	for _, w := range words {
		b := lower(w[0])
		var l map[byte]bool
		if l1[b] {
			l = l1
		} else if l2[b] {
			l = l2
		} else {
			l = l3
		}
		i := 1
		for ; i < len(w); i++ {
			b := lower(w[i])
			if !l[b] {
				break
			}
		}
		if i == len(w) {
			out = append(out, w)
		}
	}
	return out
}
