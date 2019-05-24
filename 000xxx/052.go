func totalNQueens(n int) int {
	out := 0
	cols := make([]int, n)
	for i := 0; i < n; i++ {
		cols[i] = -1
	}
	var cal func(int)
	cal = func(row int) {
		if row == n {
			out++
			return
		}
		for i := 0; i < n; i++ {
			isValid := true
			for j := 0; j < row; j++ {
				col := cols[j]
				if i == col || row-j == i-col || row-j == col-i {
					isValid = false
					break
				}
			}
			if isValid {
				cols[row] = i
				cal(row + 1)
				cols[row] = -1
			}
		}
	}
	cal(0)
	return out
}
