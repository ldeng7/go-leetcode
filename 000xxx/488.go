import "math"

func remove(board string) string {
	l, j := len(board), 0
	for i := 0; i <= l; i++ {
		if i != l && board[i] == board[j] {
			continue
		} else if i-j >= 3 {
			return remove(board[:j] + board[i:])
		}
		j = i
	}
	return board
}

func findMinStep(board string, hand string) int {
	m := [26]int{}
	for i := 0; i < len(hand); i++ {
		m[hand[i]-'A']++
	}

	var cal func(string) int
	cal = func(board string) int {
		board = remove(board)
		l := len(board)
		if 0 == l {
			return 0
		}
		c, j := math.MaxInt64, 0
		for i := 0; i <= l; i++ {
			if i != l && board[i] == board[j] {
				continue
			}
			k, b := j-i+3, board[j]-'A'
			if m[b] >= k {
				m[b] -= k
				c1 := cal(board[:j] + board[i:])
				if c1 != math.MaxInt64 && c1+k < c {
					c = c1 + k
				}
				m[b] += k
			}
			j = i
		}
		return c
	}

	o := cal(board)
	if o == math.MaxInt64 {
		return -1
	}
	return o
}
