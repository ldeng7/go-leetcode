func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func flipLights(n int, m int) int {
	n = min(n, 3)
	return min(1<<uint(n), 1+m*n)
}
