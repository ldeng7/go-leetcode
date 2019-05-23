func solve(board [][]byte) {
	if 0 == len(board) || 0 == len(board[0]) {
		return
	}
	m, n := len(board), len(board[0])
	var save func(int, int)
	save = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
			return
		}
		board[i][j] = 'S'
		save(i-1, j)
		save(i, j-1)
		save(i+1, j)
		save(i, j+1)
	}
	for i := 0; i < n; i++ {
		save(0, i)
		save(m-1, i)
	}
	for i := 1; i < m-1; i++ {
		save(i, 0)
		save(i, n-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if 'O' == board[i][j] {
				board[i][j] = 'X'
			} else if 'S' == board[i][j] {
				board[i][j] = 'O'
			}
		}
	}
}
