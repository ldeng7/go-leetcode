func numSub(s string) int {
	o, l, c := 0, len(s), 0
	ar := make([]int, l)
	for i := 0; i < l; i++ {
		if s[i] == '1' {
			c++
		} else {
			c = 0
		}
		ar[i] = c
	}
	for _, n := range ar {
		o += n
	}
	return o % 1000000007
}
