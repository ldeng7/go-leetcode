func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func minKnightMoves(x int, y int) int {
	x, y = abs(x), abs(y)
	if x+y == 1 {
		return 3
	}
	o := max(max((x+1)/2, (y+1)/2), (x+y+2)/3)
	o += (o ^ x ^ y) & 1
	return o
}
