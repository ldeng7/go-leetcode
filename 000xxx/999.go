var dirs = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func numRookCaptures(board [][]byte) int {
	findR := func() (int, int) {
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				if board[i][j] == 'R' {
					return i, j
				}
			}
		}
		return 0, 0
	}
	y0, x0 := findR()
	o := 0
	for _, d := range dirs {
		y, x := y0+d[0], x0+d[1]
		for y >= 0 && y < 8 && x >= 0 && x < 8 {
			if board[y][x] != '.' {
				if board[y][x] == 'p' {
					o++
				}
				break
			}
			y, x = y+d[0], x+d[1]
		}
	}
	return o
}
