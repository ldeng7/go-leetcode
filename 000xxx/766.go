func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 1 || len(matrix[0]) == 1 {
		return true
	}
	p := matrix[0]
	for i := 1; i < len(matrix); i++ {
		r := matrix[i]
		for j := 0; j < len(p)-1; j++ {
			if p[j] != r[j+1] {
				return false
			}
		}
		p = r
	}
	return true
}
