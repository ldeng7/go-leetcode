func rotate(matrix [][]int) {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return
	}
	l := len(matrix)
	for i := 0; i <= l<<1; i++ {
		for j := i; j < l-i-1; j++ {
			matrix[i][j], matrix[j][l-i-1], matrix[l-i-1][l-j-1], matrix[l-j-1][i] =
				matrix[l-j-1][i], matrix[i][j], matrix[j][l-i-1], matrix[l-i-1][l-j-1]
		}
	}
}
