func cal(n, k int) int {
	if n == 0 {
		return n
	}
	return (cal(n-1, k) + k) % n
}

func findTheWinner(n int, k int) int {
	return cal(n, k) + 1
}
