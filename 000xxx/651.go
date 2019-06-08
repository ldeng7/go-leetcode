func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func maxA(N int) int {
	out := N
	for i := 1; i < N-2; i++ {
		out = max(out, maxA(i)*(N-1-i))
	}
	return out
}
