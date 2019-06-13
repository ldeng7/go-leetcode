func validTicTacToe(board []string) bool {
	xwin, owin := false, false
	diag, antiDiag, iTurn := 0, 0, 0
	rows, cols := [3]int{}, [3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 'X' {
				rows[i], cols[j], iTurn = rows[i]+1, cols[j]+1, iTurn+1
				if i == j {
					diag++
				}
				if i+j == 2 {
					antiDiag++
				}
			} else if board[i][j] == 'O' {
				rows[i], cols[j], iTurn = rows[i]-1, cols[j]-1, iTurn-1
				if i == j {
					diag--
				}
				if i+j == 2 {
					antiDiag--
				}
			}
		}
	}
	xwin = rows[0] == 3 || rows[1] == 3 || rows[2] == 3 ||
		cols[0] == 3 || cols[1] == 3 || cols[2] == 3 ||
		diag == 3 || antiDiag == 3
	owin = rows[0] == -3 || rows[1] == -3 || rows[2] == -3 ||
		cols[0] == -3 || cols[1] == -3 || cols[2] == -3 ||
		diag == -3 || antiDiag == -3
	if (xwin && iTurn == 0) || (owin && iTurn == 1) {
		return false
	}
	return (iTurn == 0 || iTurn == 1) && (!xwin || !owin)
}
