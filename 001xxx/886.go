func findRotation(mat [][]int, target [][]int) bool {
	b0, b1, b2, b3 := true, true, true, true
	l := len(mat)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if mat[i][j] != target[i][j] {
				b0 = false
			}
			if mat[i][j] != target[l-j-1][i] {
				b1 = false
			}
			if mat[i][j] != target[l-i-1][l-j-1] {
				b2 = false
			}
			if mat[i][j] != target[j][l-i-1] {
				b3 = false
			}
		}
	}
	return b0 || b1 || b2 || b3
}
