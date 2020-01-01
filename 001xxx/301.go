var dirs = [3][2]int{{0, -1}, {-1, 0}, {-1, -1}}

func pathsWithMaxScore(board []string) []int {
	l := len(board)
	bss := make([][]byte, l)
	for i := 0; i < l; i++ {
		bss[i] = []byte(board[i])
	}
	bss[l-1][l-1] = '0'
	sl := l * l
	t, u := make([]int, sl), make([]int, sl)
	for i := 1; i < sl; i++ {
		t[i] = -1
	}
	u[0] = 1

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if 'X' == bss[i][j] {
				continue
			}
			c := int(bss[i][j] - '0')
			k := i*l + j
			for _, dir := range dirs {
				i1, j1 := i+dir[0], j+dir[1]
				k1 := i1*l + j1
				if i1 < 0 || j1 < 0 || t[k1] < 0 {
					continue
				}
				tc, tc1 := t[k], t[k1]+c
				if tc < tc1 {
					t[k], u[k] = tc1, u[k1]
				} else if tc == tc1 {
					u[k] = (u[k] + u[k1]) % 1000000007
				}
			}
		}
	}

	sl--
	if t[sl] < 0 {
		return []int{0, 0}
	}
	return []int{t[sl], u[sl]}
}
