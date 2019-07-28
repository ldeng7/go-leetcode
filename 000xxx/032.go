func longestValidParentheses(s string) int {
	o, l, st := 0, len(s), []int{-1}
	for i := 0; i < l; i++ {
		if s[i] == '(' {
			st = append(st, i)
		} else {
			st = st[:len(st)-1]
			if 0 == len(st) {
				st = append(st, i)
			} else if t := i - st[len(st)-1]; t > o {
				o = t
			}
		}
	}
	return o
}
