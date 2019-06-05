func countBits(num int) []int {
	out := make([]int, num+1)
	for i := 1; i <= num; i++ {
		out[i] = out[i&(i-1)] + 1
	}
	return out
}
