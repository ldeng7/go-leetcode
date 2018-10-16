func isValidSudoku(board [][]byte) bool {
	rowFlags := [9][9]bool{}
	colFlags := [9][9]bool{}
	boxFlags := [9][9]bool{}
	for i, row := range board {
		for j, b := range row {
			if '.' == b {
				continue
			}
			n := b - '1'
			if rowFlags[i][n] {
				return false
			}
			rowFlags[i][n] = true
			if colFlags[j][n] {
				return false
			}
			colFlags[j][n] = true
			if boxFlags[i/3*3+j/3][n] {
				return false
			}
			boxFlags[i/3*3+j/3][n] = true
		}
	}
	return true
}
