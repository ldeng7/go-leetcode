func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func checkZeroOnes(s string) bool {
	cal := func(b byte) int {
		o, c := -1, 0
		for i := 0; i < len(s); i++ {
			if s[i] == b {
				c++
			} else {
				c = 0
			}
			o = max(o, c)
		}
		return o
	}
	return cal('1') > cal('0')
}
