func gameOfLife(board [][]int) {
	if 0 == len(board) || 0 == len(board[0]) {
		return
	}
	m, n := len(board), len(board[0])
	dx := []int{-1, -1, -1, 0, 1, 1, 1, 0}
	dy := []int{-1, 0, 1, 1, 1, 0, -1, -1}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := 0
			for k := 0; k < 8; k++ {
				y, x := i+dx[k], j+dy[k]
				if y >= 0 && y < m && x >= 0 && x < n && (board[y][x] == 1 || board[y][x] == 2) {
					c++
				}
			}
			if board[i][j] != 0 && (c < 2 || c > 3) {
				board[i][j] = 2
			} else if board[i][j] == 0 && c == 3 {
				board[i][j] = 3
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] %= 2
		}
	}
}
