type NumMatrix struct {
	t [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return NumMatrix{nil}
	}
	m, n := len(matrix), len(matrix[0])
	t := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		t[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			t[i][j] = t[i][j-1] + t[i-1][j] - t[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{t}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if nil == this.t {
		return 0
	}
	return this.t[row2+1][col2+1] - this.t[row1][col2+1] - this.t[row2+1][col1] + this.t[row1][col1]
}
