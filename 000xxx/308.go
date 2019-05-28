type NumMatrix struct {
	matrix [][]int
	colSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return NumMatrix{}
	}
	m, n := len(matrix), len(matrix[0])
	this := &NumMatrix{matrix: matrix}
	this.colSum = make([][]int, m+1)
	this.colSum[0] = make([]int, n)
	for i := 1; i <= m; i++ {
		this.colSum[i] = make([]int, n)
		cr, crp, mrp := this.colSum[i], this.colSum[i-1], this.matrix[i-1]
		for j := 0; j < n; j++ {
			cr[j] = crp[j] + mrp[j]
		}
	}
	return *this
}

func (this *NumMatrix) Update(row int, col int, val int) {
	for i := row + 1; i < len(this.colSum); i++ {
		this.colSum[i][col] += val - this.matrix[row][col]
	}
	this.matrix[row][col] = val
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	s := 0
	for i := col1; i <= col2; i++ {
		s += this.colSum[row2+1][i] - this.colSum[row1][i]
	}
	return s
}
