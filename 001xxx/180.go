var m [1001]int

func sum(i int) int {
	if o := m[i]; o != 0 {
		return o
	}
	o := ((i + 1) * i) >> 1
	m[i] = o
	return o
}

func countLetters(S string) int {
	o, c, b := 0, 1, S[0]
	for i := 1; i < len(S); i++ {
		if S[i] == b {
			c++
		} else {
			b = S[i]
			o += sum(c)
			c = 1
		}
	}
	o += sum(c)
	return o
}
