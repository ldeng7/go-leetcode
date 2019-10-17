var dirs = [8][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	board := [8][8]bool{}
	for _, q := range queens {
		board[q[0]][q[1]] = true
	}
	o := [][]int{}
	for _, d := range dirs {
		i, j := king[0]+d[0], king[1]+d[1]
		for i >= 0 && i < 8 && j >= 0 && j < 8 {
			if board[i][j] {
				o = append(o, []int{i, j})
				break
			}
			i, j = i+d[0], j+d[1]
		}
	}
	return o
}
