func diagonalSum(mat [][]int) int {
	o, l := 0, len(mat)
	for i := 0; i < l; i++ {
		o += mat[i][i]
	}
	for i := 0; i < l; i++ {
		o += mat[i][l-i-1]
	}
	if l&1 == 1 {
		i := l >> 1
		o -= mat[i][i]
	}
	return o
}
