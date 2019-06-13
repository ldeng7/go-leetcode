func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	l := R * C
	out := make([][]int, l)
	out[0] = []int{r0, c0}
	dy, dx, i, ke := 0, 1, 1, 1
	for j := 1; i < l; j++ {
		for k := 0; k < ke; k++ {
			r0, c0 = r0+dy, c0+dx
			if r0 >= 0 && r0 < R && c0 >= 0 && c0 < C {
				out[i], i = []int{r0, c0}, i+1
			}
		}
		if j&1 == 0 {
			ke++
		}
		dy, dx = dx, -dy
	}
	return out
}
