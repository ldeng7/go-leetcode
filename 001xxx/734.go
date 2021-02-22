func decode(encoded []int) []int {
	l := len(encoded)
	a := 0
	if l&3 == 0 {
		a = 1
	}
	b := 0
	for i := 1; i < l; i += 2 {
		b ^= encoded[i]
	}
	o := make([]int, l+1)
	o[0] = a ^ b
	for i := 0; i < l; i++ {
		o[i+1] = o[i] ^ encoded[i]
	}
	return o
}
