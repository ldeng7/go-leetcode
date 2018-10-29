func solveSudoku(board [][]byte) {
	if len(board) != 9 || len(board[0]) != 9 {
		return
	}
	check := func(i, j int) bool {
		for k := 0; k < 9; k++ {
			if k != j && board[i][j] == board[i][k] {
				return false
			}
		}
		for k := 0; k < 9; k++ {
			if k != i && board[i][j] == board[k][j] {
				return false
			}
		}
		for k := i / 3 * 3; k < i/3*3+3; k++ {
			for l := j / 3 * 3; l < j/3*3+3; l++ {
				if (k != i || l != j) && board[i][j] == board[k][l] {
					return false
				}
			}
		}
		return true
	}
	var solve func(int, int) bool
	solve = func(i, j int) bool {
		if i == 9 {
			return true
		}
		if j >= 9 {
			return solve(i+1, 0)
		}
		if board[i][j] == '.' {
			for k := 1; k <= 9; k++ {
				board[i][j] = '0' + byte(k)
				if (check(i, j)) && solve(i, j+1) {
					return true
				}
				board[i][j] = '.'
			}
		} else {
			return solve(i, j+1)
		}
		return false
	}
	solve(0, 0)
}
