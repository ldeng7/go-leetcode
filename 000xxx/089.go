func grayCode(n int) []int {
	var l uint = 1 << uint(n)
	out := make([]int, l)
	for i := uint(0); i < l; i++ {
		out[i] = int((i >> 1) ^ i)
	}
	return out
}
