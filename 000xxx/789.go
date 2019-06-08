func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func escapeGhosts(ghosts [][]int, target []int) bool {
	d := abs(target[0]) + abs(target[1])
	for _, g := range ghosts {
		if d >= abs(g[0]-target[0])+abs(g[1]-target[1]) {
			return false
		}
	}
	return true
}
