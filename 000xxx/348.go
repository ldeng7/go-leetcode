type TicTacToe struct {
	n, diag, rev int
	rows, cols   []int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{n, 0, 0, make([]int, n), make([]int, n)}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	add := -1
	if 1 == player {
		add = 1
	}
	this.rows[row] += add
	this.cols[col] += add
	if row == col {
		this.diag += add
	}
	if row == this.n-col-1 {
		this.rev += add
	}
	if abs(this.rows[row]) == this.n || abs(this.cols[col]) == this.n ||
		abs(this.diag) == this.n || abs(this.rev) == this.n {
		return player
	}
	return 0
}
