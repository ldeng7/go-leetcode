func numSpecial(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	rows, cols := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				rows[i]++
				cols[j]++
			}
		}
	}
	o := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 && rows[i] == 1 && cols[j] == 1 {
				o++
			}
		}
	}
	return o
}
