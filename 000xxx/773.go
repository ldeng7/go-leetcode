const (
	tar  = 42792
	tLen = 181897
)

var moveTbl = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {3, 1, 5}, {4, 2}}

func slidingPuzzle(board [][]int) int {
	line := [6]int{board[0][0], board[0][1], board[0][2], board[1][0], board[1][1], board[1][2]}
	c := 0
	for i := 1; i < 6; i++ {
		v := line[i]
		if v == 0 {
			continue
		}
		for j := 0; j < i; j++ {
			if line[j] > v {
				c++
			}
		}
	}
	if c&1 == 1 {
		return -1
	}

	t, b := make([]bool, tLen), 0
	for _, v := range line {
		b = (b << 3) | v
	}
	if b == tar {
		return 0
	}

	ar, i := []int{b}, 1
	t[b] = true
	for 0 != len(ar) {
		ar1 := []int{}
		for _, v := range ar {
			u, v1 := 0, v
			for ; v1&0x07 != 0; u++ {
				v1 >>= 3
			}
			u = 5 - u
			for _, w := range moveTbl[u] {
				y := uint(5-w) * 3
				x := (v & ^(0x07 << y)) | (((v >> y) & 0x07) << uint((5-u)*3))
				if x == tar {
					return i
				}
				if t[x] {
					continue
				}
				t[x], ar1 = true, append(ar1, x)
			}
		}
		ar, i = ar1, i+1
	}
	return -1
}
