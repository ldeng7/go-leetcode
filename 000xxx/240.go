func searchMatrix(matrix [][]int, target int) bool {
	if 0 == len(matrix) || 0 == len(matrix[0]) ||
		matrix[0][0] > target || matrix[len(matrix)-1][len(matrix[0])-1] < target {
		return false
	}
	i, j := len(matrix)-1, 0
	for {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
		if i < 0 || j >= len(matrix[0]) {
			break
		}
	}
	return false
}
