func solve(board [][]byte) {
	if 0 == len(board) || 0 == len(board[0]) {
		return
	}
	m, n := len(board), len(board[0])
	v := make([][]bool, m)
	for i := 0; i < m; i++ {
		v[i] = make([]bool, n)
	}
	var cal func(int, int)
	cal = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || v[i][j] || board[i][j] == 'X' {
			return
		}
		v[i][j] = true
		cal(i-1, j)
		cal(i, j-1)
		cal(i+1, j)
		cal(i, j+1)
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' && !v[i][0] {
			cal(i, 0)
		}
		if board[i][n-1] == 'O' && !v[i][n-1] {
			cal(i, n-1)
		}
	}
	for j := 1; j < n-1; j++ {
		if board[0][j] == 'O' && !v[0][j] {
			cal(0, j)
		}
		if board[m-1][j] == 'O' && !v[m-1][j] {
			cal(m-1, j)
		}
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if !v[i][j] && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
