func setZeroes(matrix [][]int) {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return
	}
	m, n := len(matrix), len(matrix[0])
	rows, cols := map[int]bool{}, map[int]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i], cols[j] = true, true
			}
		}
	}
	for row, _ := range rows {
		for j := 0; j < n; j++ {
			matrix[row][j] = 0
		}
	}
	for col, _ := range cols {
		for i := 0; i < m; i++ {
			matrix[i][col] = 0
		}
	}
}
