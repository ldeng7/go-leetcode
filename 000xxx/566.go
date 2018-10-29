func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}
	out := make([][]int, r)
	for i := 0; i < r; i++ {
		out[i] = make([]int, c)
	}
	for i := 0; i < r*c; i++ {
		out[i/c][i%c] = nums[i/n][i%n]
	}
	return out
}
