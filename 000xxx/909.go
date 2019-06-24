import "math"

func snakesAndLadders(board [][]int) int {
	m, n := len(board), len(board[0])
	t := make([][]int, m)
	for i := 0; i < m; i++ {
		t[i] = make([]int, n)
	}
	t[m-1][0] = 1
	o := math.MaxInt64

	indices := func(i int) (int, int) {
		y := m - 1 - (i-1)/m
		if (m-1-y)&1 == 0 {
			return y, (i%m + m - 1) % m
		}
		return y, (m - i%m) % m
	}
	var cal func(int, int)
	cal = func(i, j int) {
		if j >= o {
			return
		} else if i >= m*m {
			if o > j {
				o = j
			}
			return
		}
		y, x := indices(i)
		if board[y][x] != -1 {
			i = board[y][x]
		}
		if i >= m*m {
			if o > j {
				o = j
			}
			return
		}
		y, x = indices(i)
		if t[y][x] > 0 && t[y][x] <= j {
			return
		}
		t[y][x] = j
		for k := 1; k <= 6; k++ {
			cal(i+k, j+1)
		}
	}

	cal(1, 0)
	if o == math.MaxInt64 {
		return -1
	}
	return o
}
