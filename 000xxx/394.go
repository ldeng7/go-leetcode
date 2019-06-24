func decodeString(s string) string {
	bs, c := []byte{}, 0
	sn, ss := []int{}, [][]byte{}
	for i := 0; i < len(s); i++ {
		b := s[i]
		if b >= '0' && b <= '9' {
			c = 10*c + int(b-'0')
		} else if b == '[' {
			sn, ss = append(sn, c), append(ss, bs)
			c, bs = 0, []byte{}
		} else if b == ']' {
			k, bst := sn[len(sn)-1], ss[len(ss)-1]
			for j := 0; j < k; j++ {
				bst = append(bst, bs...)
			}
			bs, sn, ss = bst, sn[:len(sn)-1], ss[:len(ss)-1]
		} else {
			bs = append(bs, b)
		}
	}
	if 0 != len(ss) {
		return string(ss[len(ss)-1])
	}
	return string(bs)
}
