func longestValidParentheses(s string) int {
	bs := []byte(s)
	out := 0
	for i := 0; i < len(bs)-1; i++ {
		c := bs[i]
		if ')' == c {
			continue
		}
		st := []byte{c}
		sub := []byte{c}
		for _, c1 := range bs[i+1:] {
			if c1 == '(' {
				st = append(st, c1)
				sub = append(sub, c1)
			} else {
				if 0 == len(st) {
					cnt := len(sub)
					if cnt > out {
						out = cnt
					}
					break
				}
				st = st[:len(st)-1]
				sub = append(sub, c1)
			}
			if 0 == len(st) {
				cnt := len(sub)
				if cnt > out {
					out = cnt
				}
			}
		}
	}
	return out
}
