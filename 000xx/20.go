var m = map[byte]byte{
	'}': '{',
	']': '[',
	')': '(',
}

func isValid(s string) bool {
	bs := []byte{}
	for _, c := range []byte(s) {
		switch c {
		case '{', '(', '[':
			bs = append(bs, c)
		case '}', ')', ']':
			if 0 == len(bs) {
				return false
			}
			if m[c] != bs[len(bs)-1] {
				return false
			}
			bs = bs[:len(bs)-1]
		}
	}
	return 0 == len(bs)
}
