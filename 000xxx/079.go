func exist(board [][]byte, word string) bool {
	if 0 == len(word) || 0 == len(board) || 0 == len(board[0]) {
		return false
	}
	m, n := len(board), len(board[0])
	bs := []byte(word)
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	var cal func(int, int, int) bool
	cal = func(y, x, i int) bool {
		if i == len(word) {
			return true
		}
		if y < 0 || y >= m || x < 0 || x >= n || used[y][x] || board[y][x] != bs[i] {
			return false
		}
		used[y][x] = true
		out := cal(y+1, x, i+1) || cal(y, x+1, i+1) || cal(y-1, x, i+1) || cal(y, x-1, i+1)
		used[y][x] = false
		return out
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ok := cal(i, j, 0); ok {
				return true
			}
		}
	}
	return false
}
