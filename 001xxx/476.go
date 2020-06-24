type SubrectangleQueries struct {
	r [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	r := this.r
	for i := row1; i <= row2; i++ {
		l := r[i]
		for j := col1; j <= col2; j++ {
			l[j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.r[row][col]
}
