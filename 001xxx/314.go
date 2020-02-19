func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])
	mat1 := make([][]int, m)
	for i := 0; i < m; i++ {
		mat1[i] = make([]int, n)
		copy(mat1[i], mat[i])
	}
	for i := 1; i < m; i++ {
		r, rp := mat1[i], mat1[i-1]
		for j := 0; j < n; j++ {
			r[j] += rp[j]
		}
	}
	for i := 0; i < m; i++ {
		r := mat1[i]
		for j := 1; j < n; j++ {
			r[j] += r[j-1]
		}
	}

	o := make([][]int, m)
	for i := 0; i < m; i++ {
		ar := make([]int, n)
		for j := 0; j < n; j++ {
			i1, j1, i2, j2 := max(i-k, 0), max(j-k, 0), min(i+k, m-1), min(j+k, n-1)
			v := mat1[i2][j2]
			if i1 > 0 {
				v -= mat1[i1-1][j2]
			}
			if j1 > 0 {
				v -= mat1[i2][j1-1]
			}
			if i1 > 0 && j1 > 0 {
				v += mat1[i1-1][j1-1]
			}
			ar[j] = v
		}
		o[i] = ar
	}
	return o
}
