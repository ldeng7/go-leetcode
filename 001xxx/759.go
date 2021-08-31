func countHomogenous(s string) int {
	var b1 byte = '1'
	o, c := 0, 1
	for i := 0; i < len(s); i++ {
		if b := s[i]; b == b1 {
			c++
		} else {
			b1, c = b, 1
		}
		o = (o + c) % 1000000007
	}
	return o
}
