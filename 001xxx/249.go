func minRemoveToMakeValid(s string) string {
	l, bs := len(s), []byte(s)
	st := make([]int, 0, l)
	for i := 0; i < l; i++ {
		switch s[i] {
		case '(':
			st = append(st, i)
		case ')':
			if 0 != len(st) {
				st = st[:len(st)-1]
			} else {
				bs[i] = 0
			}
		}
	}

	for 0 != len(st) {
		l := len(st) - 1
		bs[st[l]] = 0
		st = st[:l]
	}
	j := 0
	for i := 0; i < l; i++ {
		if bs[i] != 0 {
			if i != j {
				bs[j] = bs[i]
			}
			j++
		}
	}
	return string(bs[:j])
}
