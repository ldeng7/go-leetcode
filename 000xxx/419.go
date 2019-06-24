func countBattleships(board [][]byte) int {
	if 0 == len(board) || 0 == len(board[0]) {
		return 0
	}
	o, m, n := 0, len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != '.' &&
				(i == 0 || board[i-1][j] != 'X') &&
				(j == 0 || board[i][j-1] != 'X') {
				o++
			}
		}
	}
	return o
}
