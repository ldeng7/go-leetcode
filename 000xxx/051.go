import "strings"

func solveNQueens(n int) [][]string {
	out := [][]string{}
	cols := make([]int, n)
	for i := 0; i < n; i++ {
		cols[i] = -1
	}
	var cal func(int)
	cal = func(row int) {
		if row == n {
			strs := make([]string, n)
			for i := 0; i < n; i++ {
				col := cols[i]
				strs[i] = strings.Repeat(".", col) + "Q" + strings.Repeat(".", n-col-1)
			}
			out = append(out, strs)
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
