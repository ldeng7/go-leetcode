func tictactoe(moves [][]int) string {
	a, b := [8]int{}, [8]int{}
	for i, m := range moves {
		r, c := m[0], m[1]
		p := &a
		if i&1 == 1 {
			p = &b
		}
		p[r]++
		p[c+3]++
		if r == c {
			p[6]++
		}
		if r == 2-c {
			p[7]++
		}
	}
	for i := 0; i < 8; i++ {
		if a[i] == 3 {
			return "A"
		}
		if b[i] == 3 {
			return "B"
		}
	}
	if 9 == len(moves) {
		return "Draw"
	}
	return "Pending"
}
