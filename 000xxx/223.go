func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	a := (C-A)*(D-B) + (G-E)*(H-F)
	xl, yl, xr, yr := max(A, E), max(B, F), min(C, G), min(D, H)
	if xl >= xr || yl >= yr {
		return a
	}
	return a - (yr-yl)*(xr-xl)
}
