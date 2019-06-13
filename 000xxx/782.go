func movesToChessboard(board [][]int) int {
	l, sr, sc, dr, dc := len(board), 0, 0, 0, 0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if board[0][0]^board[i][0]^board[0][j]^board[i][j] == 1 {
				return -1
			}
		}
	}
	for i := 0; i < l; i++ {
		sr += board[0][i]
		sc += board[i][0]
		if board[i][0] == i&1 {
			dr++
		}
		if board[0][i] == i&1 {
			dc++
		}
	}
	if l>>1 > sr || l>>1 > sc || sr > (l+1)>>1 || sc > (l+1)>>1 {
		return -1
	}
	if l&1 == 1 {
		if dr&1 == 1 {
			dr = l - dr
		}
		if dc&1 == 1 {
			dc = l - dc
		}
	} else {
		if l-dr < dr {
			dr = l - dr
		}
		if l-dc < dc {
			dc = l - dc
		}
	}
	return (dr + dc) >> 1
}
