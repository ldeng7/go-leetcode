func updateBoard(board [][]byte, click []int) [][]byte {
	if 0 == len(board) || 0 == len(board[0]) {
		return [][]byte{}
	}
	m, n := len(board), len(board[0])
	q := [][2]int{{click[0], click[1]}}
	for 0 != len(q) {
		r, c := q[0][0], q[0][1]
		q = q[1:]
		ar, cnt := make([][2]int, 0, 8), byte(0)
		if board[r][c] == 'M' {
			board[r][c] = 'X'
		} else {
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					y, x := r+i, c+j
					if y < 0 || y >= m || x < 0 || x >= n {
						continue
					} else if board[y][x] == 'M' {
						cnt++
					} else if cnt == 0 && board[y][x] == 'E' {
						ar = append(ar, [2]int{y, x})
					}
				}
			}
		}
		if cnt > 0 {
			board[r][c] = cnt + '0'
		} else {
			for _, p := range ar {
				board[p[0]][p[1]] = 'B'
				q = append(q, p)
			}
		}
	}
	return board
}
