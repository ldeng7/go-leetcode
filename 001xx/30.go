func protect(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'T'
	protect(board, i-1, j)
	protect(board, i, j-1)
	protect(board, i+1, j)
	protect(board, i, j+1)
}

func solve(board [][]byte) {
	if 0 == len(board) || 0 == len(board[0]) {
		return
	}
	m, n := len(board), len(board[0])
	for i := 0; i < n; i++ {
		protect(board, 0, i)
		protect(board, m-1, i)
	}
	for i := 1; i < m-1; i++ {
		protect(board, i, 0)
		protect(board, i, n-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if 'O' == board[i][j] {
				board[i][j] = 'X'
			} else if 'T' == board[i][j] {
				board[i][j] = 'O'
			}
		}
	}
}
