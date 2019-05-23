func numMovesStones(a int, b int, c int) []int {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	min, max := 2, c-a-2
	if c-a == 2 {
		min = 0
	} else if b-a <= 2 || c-b <= 2 {
		min = 1
	}
	return []int{min, max}
}
