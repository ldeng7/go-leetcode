func decode(encoded []int, first int) []int {
	l := len(encoded)
	o := make([]int, l+1)
	o[0] = first
	for i := 0; i < l; i++ {
		o[i+1] = o[i] ^ encoded[i]
	}
	return o
}
