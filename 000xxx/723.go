func candyCrush(board [][]int) [][]int {
	m, n := len(board), len(board[0])
	for {
		ds := [][2]int{}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if board[i][j] == 0 {
					continue
				}
				y0, x0, y1, x1 := i, j, i, j
				for y0 >= 0 && y0 > i-3 && board[y0][j] == board[i][j] {
					y0--
				}
				for y1 < m && y1 < i+3 && board[y1][j] == board[i][j] {
					y1++
				}
				for x0 >= 0 && x0 > j-3 && board[i][x0] == board[i][j] {
					x0--
				}
				for x1 < n && x1 < j+3 && board[i][x1] == board[i][j] {
					x1++
				}
				if y1-y0 > 3 || x1-x0 > 3 {
					ds = append(ds, [2]int{i, j})
				}
			}
		}
		if 0 == len(ds) {
			break
		}
		for _, d := range ds {
			board[d[0]][d[1]] = 0
		}
		for j := 0; j < n; j++ {
			k := m - 1
			for i := m - 1; i >= 0; i-- {
				if board[i][j] != 0 {
					board[k][j], board[i][j] = board[i][j], board[k][j]
					k--
				}
			}
		}
	}
	return board
}
