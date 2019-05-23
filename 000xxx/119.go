func getRow(rowIndex int) []int {
	up := []int{1}
	for i := 1; i <= rowIndex; i++ {
		row := make([]int, i+1)
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < len(row)-1; j++ {
			row[j] = up[j-1] + up[j]
		}
		up = row
	}
	return up
}
