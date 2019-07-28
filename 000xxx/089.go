func grayCode(n int) []int {
	var l uint = 1 << uint(n)
	o := make([]int, l)
	for i := uint(0); i < l; i++ {
		o[i] = int((i >> 1) ^ i)
	}
	return o
}
