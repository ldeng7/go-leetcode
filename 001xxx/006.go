func special(n int) int {
	if n == 3 {
		return 6
	}
	return n
}

func clumsy(N int) int {
	if N <= 3 {
		return special(N)
	}
	i, ie, out := N, N&3, 0
	if i >= 4 {
		out = i*(i-1)/(i-2) + (i - 3)
		i -= 4
	}
	for ; i != ie; i -= 4 {
		out += -(i * (i - 1) / (i - 2)) + (i - 3)
	}
	return out - special(ie)
}
