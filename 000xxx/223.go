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
	l := max(A, E)
	r := max(min(C, G), l)
	b := max(B, F)
	t := max(min(D, H), b)
	return (C-A)*(D-B) - (r-l)*(t-b) + (G-E)*(H-F)
}
