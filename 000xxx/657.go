func judgeCircle(moves string) bool {
	c1, c2 := 0, 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'L':
			c1++
		case 'R':
			c1--
		case 'U':
			c2++
		case 'D':
			c2--
		}
	}
	return c1 == 0 && c2 == 0
}
