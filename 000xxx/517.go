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

func findMinMoves(machines []int) int {
	s := 0
	for _, m := range machines {
		s += m
	}
	if s%len(machines) != 0 {
		return -1
	}
	o, c, a := 0, 0, s/len(machines)
	for _, m := range machines {
		c += m - a
		o = max(o, max(abs(c), m-a))
	}
	return o
}
